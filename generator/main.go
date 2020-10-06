package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"go/format"

	"github.com/stoewer/go-strcase"
)

func main() {

	dir, err := os.Getwd()

	if filepath.Base(dir) == "generator" {
		dir = filepath.Join(dir, "..")
	}

	data, err := downloadSwaggerFile()
	if err != nil {
		panic(err)
	}
	os.MkdirAll(filepath.Join(dir, "api", "models"), 0755)
	os.MkdirAll(filepath.Join(dir, "api"), 0755)

	generateModels(dir, *data.Definitions)

	funcnames := generateApi(dir, *data.Paths)
	b, e := ioutil.ReadFile(path.Join(dir, "generator", "tmpl", "api.go"))
	if e != nil {
		panic(e)
	}
	content := strings.Replace(string(b), "// functions", strings.Join(funcnames, "\n"), 1)
	b, e = format.Source([]byte(content))
	if e != nil {
		fmt.Println(content)
		fmt.Println(e)
		os.Exit(1)
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "api", "api.go"), b, 0644); err != nil {
		panic(err)
	}
	if e := exec.Command("goimports", "-w", filepath.Join(dir, "api")).Run(); e != nil {
		panic(e)
	}
}

func generateApi(dir string, paths map[string]map[string]swaggerPayloadPath) []string {

	files := map[string][]string{}
	var funcnames []string
	var orderedendpoints []string
	for endpoint := range paths {
		orderedendpoints = append(orderedendpoints, endpoint)
	}
	sort.Strings(orderedendpoints)
	for _, endpoint := range orderedendpoints {
		methods := paths[endpoint]
		var orderedmethods []string
		for method := range methods {
			orderedmethods = append(orderedmethods, method)
		}
		sort.Strings(orderedmethods)
		for _, method := range orderedmethods {
			details := methods[method]
			var line []string
			if details.Summary == nil {
				fmt.Println("missing summery, skipping")
				continue
			}
			if details.Description != nil {
				descs := strings.Split(*details.Description, "\n")
				for _, d := range descs {
					line = append(line, "// "+d)
					funcnames = append(funcnames, "// "+d)
				}
			}
			var returntype string
			summary := strings.TrimSuffix(*details.Summary, ".")
			summary = strings.ReplaceAll(summary, "(", "")
			summary = strings.ReplaceAll(summary, ")", "")
			funcname := strings.ReplaceAll(summary, " ", "")
			if response := details.Responses; response != nil {
				resp := *response
				if good, ok := resp["200"]; ok {
					if good.Schema != nil && good.Schema.Items != nil && good.Schema.Items.Ref != nil {
						returntype = "*[]models." + strings.Replace(*good.Schema.Items.Ref, "#/definitions/", "", 1)
					}
					if good.Schema != nil && good.Schema.Ref != nil {
						returntype = "*models." + strings.Replace(*good.Schema.Ref, "#/definitions/", "", 1)
					}
				}
			}
			var funcsignature string
			var params []string
			var body []string
			if details.Parameters != nil {
				var vars []string
				for _, param := range *details.Parameters {
					shortname := strcase.LowerCamelCase(*param.Name)
					if *param.In == "path" {
						if param.Format == nil || param.Type == nil {
							fmt.Println("missing param types")
							continue
						}
						if *param.Format == "int64" {
							vars = append(vars, shortname+" int64")
							endpoint = strings.Replace(endpoint, "{"+*param.Name+"}", `"+fmt.Sprint(`+shortname+`)+"`, 1)
						} else {
							vars = append(vars, shortname+" string")
							endpoint = strings.Replace(endpoint, "{"+*param.Name+"}", `"+`+shortname+`+"`, 1)
						}
					} else if *param.In == "body" {
						if param.Schema == nil {
							fmt.Println("missing param types")
							continue
						}
						if param.Schema.Ref == nil {
							fmt.Println("missing param types")
							continue
						}
						if method == "post" || method == "put" {
							body = append(params, "if "+shortname+" != nil {")
							body = append(body, "jsonbody, _ := json.Marshal("+shortname+")")
							body = append(body, "body = bytes.NewBuffer(jsonbody)")
							body = append(body, "}")
						} else {
							params = append(params, "if "+shortname+" != nil {")
							params = append(params, "	kv:=map[string]interface{}{}")
							params = append(params, "	b,_:=json.Marshal("+shortname+")")
							params = append(params, "	json.Unmarshal(b, &kv)")
							params = append(params, "	for k, v :=range kv {")
							params = append(params, "		params.Set(k,fmt.Sprint(v));")
							params = append(params, "	}")
							params = append(params, "}")
						}
						vars = append(vars, shortname+" *models."+strings.Replace(*param.Schema.Ref, "#/definitions/", "", 1))
					}
				}
				if returntype != "" {
					funcsignature = funcname + "(" + strings.Join(vars, ",") + ") (" + returntype + ", error) "
				} else {
					funcsignature = funcname + "(" + strings.Join(vars, ",") + ") error "
				}
			} else {
				if returntype != "" {
					funcsignature = funcname + "()(" + returntype + ", error)"
				} else {
					funcsignature = funcname + "() error "
				}
			}
			funcnames = append(funcnames, funcsignature)
			line = append(line, "func (a*api)"+funcsignature+"{")
			line = append(line, "params := url.Values{}")
			if params != nil {
				line = append(line, params...)
			}
			if body != nil {
				line = append(line, "var body *bytes.Buffer")
				line = append(line, body...)
			}
			if returntype != "" {
				line = append(line, "var out "+strings.TrimPrefix(returntype, "*"))
			} else {
				line = append(line, "var out interface{}")
			}
			if body != nil {
				line = append(line, `if err:= a.request("`+strings.ToUpper(method)+`","`+endpoint+`",  params, body, &out); err != nil {`)
			} else {
				line = append(line, `if err:= a.request("`+strings.ToUpper(method)+`","`+endpoint+`",  params, nil, &out); err != nil {`)
			}
			if returntype != "" {
				line = append(line, "return nil, err")
			} else {
				line = append(line, "return err")
			}
			line = append(line, "}")
			if returntype != "" {
				line = append(line, "return &out, nil")
			} else {
				line = append(line, "return  nil")
			}
			line = append(line, "}")

			content := strings.Join(line, "\n")
			b, e := format.Source([]byte(content))
			if e != nil {
				fmt.Println(content)
				panic(e)
			}

			filename := strings.TrimPrefix(funcname, "Delete")
			filename = strings.TrimPrefix(filename, "Get")
			filename = strings.TrimPrefix(filename, "List")
			filename = strings.TrimPrefix(filename, "Update")
			filename = strings.TrimPrefix(filename, "Create")
			if strings.HasSuffix(filename, "ies") {
				filename = strings.TrimSuffix(filename, "ies")
				filename += "y"
			} else {
				filename = strings.TrimSuffix(filename, "s")
			}
			files[filename] = append(files[filename], string(b))
		}

	}
	for filename, funcs := range files {
		var header []string
		header = append(header, "package api")
		header = append(header, "import (")
		header = append(header, `"fmt"`)
		header = append(header, `"bytes"`)
		header = append(header, `"net/url"`)
		header = append(header, `"github.com/pec1985/go-clubhouse/api/models"`)
		header = append(header, `"encoding/json"`)
		header = append(header, ")")
		funcs = append(header, funcs...)
		content := strings.Join(funcs, "\n")
		b, e := format.Source([]byte(content))
		if e != nil {
			fmt.Println(content)
			fmt.Println(e)
			os.Exit(1)
		}
		ioutil.WriteFile(filepath.Join(dir, "api", filename+".go"), b, 0644)
	}
	return funcnames
}

func generateModels(dir string, definitions map[string]swaggerPayloadDefinition) {

	var ordereddefs []string
	for name := range definitions {
		ordereddefs = append(ordereddefs, name)
	}
	sort.Strings(ordereddefs)
	for _, name := range ordereddefs {
		def := definitions[name]
		var lines []string
		lines = append(lines, "package models")
		lines = append(lines, "type "+name+" struct {")
		var orderedprops []string
		for name := range *def.Properties {
			orderedprops = append(orderedprops, name)
		}
		sort.Strings(orderedprops)
		for _, n := range orderedprops {
			p := (*def.Properties)[n]
			l := generatePropery(n, p)
			if l != "" {
				lines = append(lines, l)
			}
		}
		lines = append(lines, "}")
		lines = append(lines, "func (m*"+name+") Stringify() string {")
		lines = append(lines, "b,_:=json.Marshal(m)")
		lines = append(lines, "return string(b)")
		lines = append(lines, "}")
		lines = append(lines, "func (m*"+name+") StringifyPretty() string {")
		lines = append(lines, `b,_:=json.MarshalIndent(m,"","  ")`)
		lines = append(lines, "return string(b)")
		lines = append(lines, "}")
		object := strings.Join(lines, "\n")
		b, e := format.Source([]byte(object))
		if e != nil {
			fmt.Println(object)
			panic(e)
		}
		ioutil.WriteFile(filepath.Join(dir, "api", "models", name+".go"), b, 0644)
	}
}

func generatePropery(name string, property swaggerPayloadDefinitionProperty) string {

	line := strcase.UpperCamelCase(name)
	if null := property.XNullable; null != nil && *null {
		line += "*"
	}
	line += " "
	if property.Format != nil {
		switch *property.Format {
		case "date-time":
			line += "time.Time"
		case "uuid":
			line += "string"
		case "int64":
			line += "int64"
		case "double":
			line += "float64"
		}
	} else if property.Type != nil {
		switch *property.Type {
		case "string":
			line += "string"
		case "boolean":
			line += "bool"
		case "file":
			line += "File"
		case "array":
			items := *property.Items
			if ref := items.Ref; ref != nil {
				line += "[]" + strings.Replace(*ref, "#/definitions/", "", 1)
			} else {
				if items.Format != nil {
					switch *items.Format {
					case "date-time":
						line += "[]time.Time"
					case "uuid":
						line += "[]string"
					case "int64":
						line += "[]int64"
					case "double":
						line += "[]float64"
					}
				} else {
					switch *items.Type {
					case "string":
						line += "[]string"
					case "boolean":
						line += "[]bool"
					case "file":
						line += "[]File"
					}
				}
			}
			break
		}
	} else if ref := property.Ref; ref != nil {
		line += strings.Replace(*ref, "#/definitions/", "", 1)
	} else {
		fmt.Println("could not find type")
		return ""
	}
	line += " `json:\"" + name + "\"`"
	if property.Description != nil {
		return "// " + *property.Description + "\n" + line
	}
	return line
}
