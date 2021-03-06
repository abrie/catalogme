package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Columns []Column
type Schema map[string]Columns

func toObjName(tableName string) string {
	parts := strings.Split(tableName, "_")
	for idx, part := range parts {
		if part == "id" {
			parts[idx] = "ID"
		} else {
			parts[idx] = strings.Title(part)
		}
	}

	return strings.Join(parts, "")
}

func toSelectField(name string) string {
	return name
}

func toScanParam(name string) string {
	parts := strings.Split(name, "_")
	for idx, part := range parts {
		if part == "id" {
			parts[idx] = "ID"
		} else {
			parts[idx] = strings.Title(part)
		}
	}

	return strings.Join(parts, "")
}

func toUpdateParam(name string) string {
	parts := strings.Split(name, "_")
	for idx, part := range parts {
		if part == "id" {
			parts[idx] = "ID"
		} else {
			parts[idx] = strings.Title(part)
		}
	}

	return strings.Join(parts, "")
}

func toUpdateFieldValues(columns []Column) string {
	var params []string
	for _, column := range columns {
		if column.Name != "id" {
			params = append(params, fmt.Sprintf("input.%s", toScanParam(column.Name)))
		}
	}

	return strings.Join(params, ",")
}

func toUpdateField(name string) string {
	return fmt.Sprintf(`%s=?`, name)
}

func toUpdateFieldNames(columns []Column) string {
	var params []string
	for _, column := range columns {
		if column.Name != "id" {
			params = append(params, toUpdateField(column.Name))
		}
	}

	return strings.Join(params, ",")
}

func toInsertField(name string) string {
	return name
}

func toInsertFieldNames(columns []Column) string {
	var params []string
	for _, column := range columns {
		if column.Name != "id" {
			params = append(params, toInsertField(column.Name))
		}
	}

	return strings.Join(params, ",")
}

func toInsertFieldParams(columns []Column) string {
	var params []string
	for _, column := range columns {
		if column.Name != "id" {
			params = append(params, "?")
		}
	}

	return strings.Join(params, ",")
}

func toInsertFieldValues(columns []Column) string {
	var params []string
	for _, column := range columns {
		if column.Name != "id" {
			params = append(params, fmt.Sprintf("input.%s", toScanParam(column.Name)))
		}
	}

	return strings.Join(params, ",")
}

func toSelectFieldNames(columns []Column) string {
	var params []string
	for _, column := range columns {
		params = append(params, toSelectField(column.Name))
	}

	return strings.Join(params, ",")
}

func toSelectFieldValues(columns []Column) string {
	var params []string
	for _, column := range columns {
		params = append(params, fmt.Sprintf("&obj.%s", toScanParam(column.Name)))
	}

	return strings.Join(params, ",")
}

func isForeignKey(name string) bool {
	return strings.HasSuffix(name, "_id")
}

func filterInputColumns(input []Column) []Column {
	var output []Column
	for _, column := range input {
		if column.Name == "id" {
			continue
		} else {
			output = append(output, column)
		}
	}

	return output
}

func generateGQLSchema(schema *Schema, templateName string, outputPath string) {
	type Table struct {
		ObjType   string
		TableName string
		Fields    []Column
		Inputs    []Column
	}

	tables := map[string]Table{}
	for tableName, columns := range *schema {
		table := Table{
			TableName: tableName,
			ObjType:   toObjName(tableName),
			Fields:    columns,
			Inputs:    filterInputColumns(columns),
		}

		tables[tableName] = table
	}

	/* Re-scan the columns in each table, and add *_list fields
	if the column references another table.
	*/
	for tableName, table := range tables {
		for _, column := range table.Fields {
			if strings.HasSuffix(column.Name, "_id") {
				parentName := strings.TrimSuffix(column.Name, "_id")
				parent := tables[parentName]
				newField := Column{Name: tableName + "_list", Type: fmt.Sprintf("[%s!]!", toObjName(tableName))}
				parent.Fields = append(parent.Fields, newField)
				tables[parentName] = parent
			}
		}
	}

	list := []Table{}
	for _, val := range tables {
		list = append(list, val)
	}

	funcs := template.FuncMap{
		"GetFieldType": func(column Column) string {
			if strings.HasSuffix(column.Name, "_id") {
				return "ID"
			} else {
				switch column.Name {
				case "id":
					return "ID"
				}
				switch column.Type {
				case "text":
					return "String"
				case "integer":
					return "Int"
				default:
					return column.Type
				}
			}
		},
	}
	outBytes, err := executeTemplate(templateName, list, funcs)
	if err != nil {
		log.Fatal(err)
	}

	writeBytes(outBytes.Bytes(), outputPath)
}

func generateLoaderCode(schema *Schema, templateName string, outPath string) {
	type Table struct {
		ObjType           string
		TableName         string
		SelectFieldNames  string
		SelectFieldValues string
		ForeignKey        string
		UpdateFieldNames  string
		UpdateFieldValues string
		InsertFieldNames  string
		InsertFieldParams string
		InsertFieldValues string
	}

	var tables []Table

	for tableName, columns := range *schema {
		table := Table{
			ObjType:           toObjName(tableName),
			TableName:         tableName,
			SelectFieldNames:  toSelectFieldNames(columns),
			SelectFieldValues: toSelectFieldValues(columns),
			UpdateFieldNames:  toUpdateFieldNames(columns),
			UpdateFieldValues: toUpdateFieldValues(columns),
			InsertFieldNames:  toInsertFieldNames(columns),
			InsertFieldParams: toInsertFieldParams(columns),
			InsertFieldValues: toInsertFieldValues(columns)}

		for _, column := range columns {
			if isForeignKey(column.Name) {
				table.ForeignKey = column.Name
			}
		}

		tables = append(tables, table)
	}

	outBytes, err := executeTemplate(templateName, tables, template.FuncMap{})
	if err != nil {
		log.Fatal(err)
	}

	formattedBytes, err := format.Source(outBytes.Bytes())
	if err != nil {
		writeBytes(outBytes.Bytes(), "error")
		log.Fatal("Formatter failed:", err)
	}

	writeBytes(formattedBytes, outPath)
}

func executeTemplate(name string, data interface{}, funcs template.FuncMap) (*bytes.Buffer, error) {
	templates, err := template.New("generator-templates").Funcs(funcs).ParseFiles(name)
	if err != nil {
		return nil, err
	}

	tmpl := templates.Lookup(path.Base(name))
	if tmpl == nil {
		return nil, fmt.Errorf("No template named %s", name)
	}

	var outBytes bytes.Buffer
	writer := bufio.NewWriter(&outBytes)
	err = tmpl.Execute(writer, data)
	if err != nil {
		return nil, err
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	return &outBytes, nil
}

func main() {
	var inputFile = flag.String("input", "", "JSON describing the DB tables.")
	var outputPath = flag.String("output", "", "Path to place generated files.")
	var templatesPath = flag.String("templates", "", "Path to templates.")
	flag.Parse()

	if *inputFile == "" || *outputPath == "" || *templatesPath == "" {
		flag.PrintDefaults()
		return
	}

	schema, err := loadSchema(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	generateLoaderCode(schema, path.Join(*templatesPath, "datastore.tmpl"), path.Join(*outputPath, "datastore.go"))
	generateGQLSchema(schema, path.Join(*templatesPath, "schema.graphql.tmpl"), path.Join(*outputPath, "schema.graphql"))
}

func writeBytes(bytes []byte, filename string) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		log.Printf("%s, %s\n", filename, path.Dir(filename))
		log.Fatal(err)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	outFile.Write(bytes)
}

func loadSchema(schemaPath string) (*Schema, error) {
	var (
		schema Schema
	)

	jsonFile, err := os.Open(schemaPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonBytes, &schema)

	return &schema, nil
}
