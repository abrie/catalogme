package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Columns []string
type Schema map[string]Columns

type Table struct {
	ObjType      string
	TableName    string
	SelectParams string
	ScanParams   string
}

type ListTable struct {
	Table      Table
	ForeignKey string
}

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
	if name == "id" {
		return "ID"
	} else {
		parts := strings.Split(name, "_")
		for idx, part := range parts {
			parts[idx] = strings.Title(part)
		}

		return strings.Join(parts, "")
	}
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

func main() {
	schema, err := loadSchema("../migrate/merged/merged.json")
	if err != nil {
		log.Fatal(err)
	}

	templates, err := loadTemplate()
	if err != nil {
		log.Fatal(err)
	}

	for tableName, columns := range *schema {
		table := Table{
			ObjType:      toObjName(tableName),
			TableName:    tableName,
			SelectParams: toSelectParams(columns),
			ScanParams:   toScanParams(columns)}

		err = templates.ExecuteTemplate(os.Stdout, "queryOne.tmpl", table)
		if err != nil {
			panic(err)
		}

		for _, column := range columns {
			log.Println(column)
			if isForeignKey(column) {
				params := ListTable{
					ForeignKey: column,
					Table:      table}

				err = templates.ExecuteTemplate(os.Stdout, "queryList.tmpl", params)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func loadTemplate() (*template.Template, error) {
	templates, err := template.ParseFiles("queryOne.tmpl", "queryList.tmpl")
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
