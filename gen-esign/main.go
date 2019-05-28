// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gen-esign creates the esign subpackages based upon DocuSign's
// esignature.rest.swagger.json definition file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"text/template"
)

const (
	openAPIdefinition = "https://github.com/docusign/eSign-OpenAPI-Specification"
	generatorVersion  = "20171011"
)

var (
	basePkg     = flag.String("basepkg", "github.com/jfcote87/esign", "root package in gopath")
	baseDir     = flag.String("gopath", fmt.Sprintf("%s/src", build.Default.GOPATH), "GOPATH src directory")
	templDir    = flag.String("template", "gen-esign/templates", "directory containing output templates.")
	buildFlag   = flag.Bool("build", false, "Compile generated packages.")
	swaggerFile = flag.String("swagger_file", "gen-esign/esignature.rest.swagger.json", "If non-empty, the path to a local file on disk containing the API to generate. Exclusive with setting --api.")
	skipFormat  = flag.Bool("skip_format", false, "skip gofmt command")
)

// main program
func main() {
	doc, err := getDocument()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Put the Definitions (structs) in order
	sort.Sort(doc.Definitions)
	defMap := make(map[string]Definition)
	structName := ""
	defList := make([]Definition, 0, len(doc.Definitions))
	// create defMap for lookup later field and parameter
	// lookups.  Make certain defList has only unique names.

	for _, def := range doc.Definitions {
		defMap["#/definitions/"+def.ID] = def
		if structName != def.Name {
			defList = append(defList, def)
		}
		structName = def.Name
	}
	// create templates
	if !strings.HasPrefix(*templDir, "/") {
		*templDir = path.Join(*baseDir, *basePkg, *templDir)
	}
	genTemplates, err := template.ParseFiles(path.Join(*templDir, "service.tmpl"), path.Join(*templDir, "/model.tmpl"))
	if err != nil {
		log.Fatalf("Templates: %v", err)
	}

	// generate model.go first
	if err := doModel(genTemplates.Lookup("model.tmpl"), defList, defMap); err != nil {
		log.Fatalf("Generating model.go failed: %v", err)
	}

	sort.Sort(doc.Operations)
	ops := make(map[string][]Operation, 0)
	for _, op := range doc.Operations {
		if op.Status == "restricted" {
			log.Printf("Skipping: %s %s", op.Status, op.OperationID)
			continue
		}
		if op.Service == "" {
			log.Printf("No service specified: %s", op.OperationID)
			continue
		}
		if !OperationSkipList[op.OperationID] {
			if overrideService, ok := serviceOverrides[op.OperationID]; ok {
				op.Service = overrideService
			}
			opList := ops[op.Service]
			opList = append(opList, op)
			ops[op.Service] = opList
		}
	}
	tagDescMap := make(map[string]string)
	for _, tag := range doc.DSTags {
		tagDescMap[tag.Name] = tag.Description
	}

	for k, v := range ops {
		log.Printf("Generating %s", k)
		descrip, _ := tagDescMap[k]

		if err := doPackage(genTemplates.Lookup("service.tmpl"), k, descrip, v, defMap); err != nil {
			log.Fatalf("generate %s.go failed: %v", k, err)
		}
	}
}

// getDocument loads the swagger def file and applies overrides
func getDocument() (*Document, error) {
	if err := os.Chdir(getEsignDir()); err != nil {
		log.Fatalf("unable to chdir to %s: %v", getEsignDir(), err)
	}
	// Open swagger file and parse
	f, err := os.Open(*swaggerFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to open: %v", err)
	}
	defer f.Close()
	var doc *Document
	if err = json.NewDecoder(f).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Unable to parse: %v", err)
	}
	var opAdditions OpList
	// Add additional operations from overrides package
	for i, op := range doc.Operations {
		if strings.Contains(op.Description, "**Deprecated**") {
			doc.Operations[i].Deprecated = true
		}
		// add media download when empty get response
		if op.HTTPMethod == "GET" && op.Responses["200"].Schema == nil {
			newResp := op.Responses["200"]
			newResp.Schema = &SchemaRef{
				Type: "file",
			}
			doc.Operations[i].Responses["200"] = newResp
		}
	}
	doc.Operations = append(doc.Operations, opAdditions...)

	return doc, nil
}

// doModel generates the model.go in the model package
func doModel(modelTempl *template.Template, defList []Definition, defMap map[string]Definition) error {
	// create model.go
	f, err := makePackageFile("model") //os.Create("model.go")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		if err == nil && !*skipFormat {
			exec.Command("gofmt", "-s", "-w", "model.go").Run()
		}
	}()
	// get field overrides and tab overrides
	fldOverrides := GetFieldOverrides()
	tabDefs := TabDefs(defMap, fldOverrides)
	var data = struct {
		Definitions  []Definition
		DefMap       map[string]Definition
		FldOverrides map[string]map[string]string
		CustomCode   string
	}{
		Definitions:  append(tabDefs, defList...), // Prepend tab definitions
		DefMap:       defMap,
		FldOverrides: fldOverrides,
		CustomCode:   CustomCode,
	}
	return modelTempl.Execute(f, data)
}

// ExtOperation contains all needed info
// for the template merge
type ExtOperation struct {
	Operation
	OpPayload         *Payload
	HasUploads        bool
	IsMediaUpload     bool
	PathParams        []PathParam
	FuncName          string
	QueryOptions      []QueryOpt
	Result            string
	DownloadAdditions []DownloadAddition
}

// doPackage creates a subpackage go file
func doPackage(resTempl *template.Template, serviceName, description string, ops []Operation, defMap map[string]Definition) error {
	packageName := strings.ToLower(serviceName)
	comments := strings.Split(strings.TrimRight(description, "\n"), "\n")
	if packageName == "uncategorized" {
		packageName = "future"
		comments = append(comments, "Future calls may change or move to other packages.")
	}
	f, err := makePackageFile(packageName)
	if err != nil {
		return err
	}

	extOps := make([]ExtOperation, 0, len(ops))
	paramOverrides := GetParameterOverrides()

	for _, op := range ops {
		payload := op.Payload(defMap)
		extOps = append(extOps, ExtOperation{
			Operation:         op,
			OpPayload:         payload,
			HasUploads:        IsUploadFilesOperation(op.OperationID),
			IsMediaUpload:     payload != nil && payload.Type == "*esign.UploadFile",
			PathParams:        op.PathParameters(),
			FuncName:          op.GoFuncName(GetServicePrefixes(op.Service)),
			QueryOptions:      op.QueryOpts(paramOverrides),
			Result:            op.Result(defMap),
			DownloadAdditions: GetDownloadAdditions(op.OperationID),
		})
	}
	var data = struct {
		Service    string
		Package    string
		Directory  string
		Operations []ExtOperation
		Comments   []string
		Packages   []string
	}{
		Service:    serviceName,
		Package:    packageName,
		Directory:  *basePkg,
		Operations: extOps,
		Comments:   comments,
		Packages:   []string{`"context"`, `"net/url"`},
	}
	importMap := make(map[string]bool)
	for _, op := range extOps {
		if len(op.PathParameters()) > 0 {
			importMap[`"strings"`] = true
			break
		}
	}
	for _, o := range extOps {
		for _, q := range o.QueryOptions {
			switch q.Type {
			case "...string":
				importMap[`"strings"`] = true
			case "int":
				importMap[`"fmt"`] = true
			case "float64":
				importMap[`"fmt"`] = true
			case "time.Time":
				importMap[`"time"`] = true
			}
		}
		if o.IsMediaUpload {
			importMap[`"io"`] = true
		}
	}

	for k, v := range importMap {
		if v {
			data.Packages = append(data.Packages, k)
		}
	}
	sort.Strings(data.Packages)
	data.Packages = append(data.Packages,
		"",
		"\""+*basePkg+"\"",
		"\""+*basePkg+"/model\"")

	defer func() {
		f.Close()
		if err == nil && !*skipFormat {
			exec.Command("gofmt", "-s", "-w", packageName+".go").Run()
		}
	}()
	return resTempl.Execute(f, data)
}

func getEsignDir() string {
	return path.Join(*baseDir, *basePkg)
}

func makePackageFile(packageName string) (*os.File, error) {
	pkgDir := getEsignDir() + "/" + packageName
	if err := os.Chdir(pkgDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(pkgDir, 0755); err == nil {
				os.Chdir(pkgDir)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	return os.Create(packageName + ".go")
}