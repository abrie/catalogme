package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Columns []string
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
	if name == "id" {
		return "ROWID"
	} else {
		return name
	}
}

func toScanField(name string) string {
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

func toSelectParams(columns []string) string {
	var params []string
	for _, column := range columns {
		params = append(params, toSelectField(column))
	}

	return strings.Join(params, ",")
}

func toScanParams(columns []string) string {
	var params []string
	for _, column := range columns {
		params = append(params, fmt.Sprintf("&obj.%s", toScanField(column)))
	}

	return strings.Join(params, ",")
}

func isForeignKey(name string) bool {
	return strings.HasSuffix(name, "_id")
}

func generateGQLSchema(schema *Schema) {
	type Table struct {
		ObjType string
		Fields  []string
	}

	var tables []Table
	for tableName, columns := range *schema {
		table := Table{
			ObjType: toObjName(tableName),
			Fields:  columns,
		}

		tables = append(tables, table)
	}
}

func generateLoaderCode(schema *Schema) {
	type Table struct {
		ObjType      string
		TableName    string
		SelectParams string
		ScanParams   string
		ForeignKey   string
	}

	var tables []Table

	for tableName, columns := range *schema {
		table := Table{
			ObjType:      toObjName(tableName),
			TableName:    tableName,
			SelectParams: toSelectParams(columns),
			ScanParams:   toScanParams(columns)}

		for _, column := range columns {
			if isForeignKey(column) {
				table.ForeignKey = column
			}
		}

		tables = append(tables, table)
	}

	outBytes, err := executeTemplate("datastore.tmpl", tables)
	if err != nil {
		log.Fatal(err)
	}

	formattedBytes, err := format.Source(outBytes.Bytes())
	if err != nil {
		writeBytes(outBytes.Bytes(), "datastore.bad")
		log.Fatal("Formatter failed:", err)
	}

	writeBytes(formattedBytes, "datastore.go")
}

func executeTemplate(name string, data interface{}) (*bytes.Buffer, error) {
	templates, err := template.ParseFiles(name)
	if err != nil {
		return nil, err
	}

	tmpl := templates.Lookup(name)
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
	schema, err := loadSchema("../migrate/merged/merged.json")
	if err != nil {
		log.Fatal(err)
	}

	generateLoaderCode(schema)
	generateGQLSchema(schema)
}

func writeBytes(bytes []byte, filename string) {
	outFile, err := os.Create("datastore.go")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	outFile.Write(bytes)
}

func loadSchemaTemplate() (*template.Template, error) {
	templates, err := template.ParseFiles("schema.graphql.tmpl")
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func loadCodeTemplate() (*template.Template, error) {
	templates, err := template.ParseFiles("datastore.tmpl")
	if err != nil {
		return nil, err
	}

	return templates, nil
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
