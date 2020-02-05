package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type Column map[string]string
type Table []Column

type Data struct {
	Name  string
	Table Table
}

type TemplateData struct {
	Input  string
	Tables []Data
}

type Database map[string]Table

func CreateFields(table []Column) (string, error) {
	var list []string
	for _, column := range table {
		list = append(list, fmt.Sprintf(`"%s" %s`, column["name"], column["type"]))
	}
	return strings.Join(list, ", "), nil
}

func InsertFields(table []Column) (string, error) {
	var list []string
	for _, column := range table {
		list = append(list, fmt.Sprintf(`"%s"`, column["name"]))
	}
	return strings.Join(list, ", "), nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Expected 3 parameters: <database.sqlite3> <tables.json> <reshape.sql.gotmpl>")
		return
	}

	databaseFile := os.Args[1]
	tablesJSON := os.Args[2]
	templateFile := os.Args[3]

	database := Database{}
	reader, err := os.Open(tablesJSON)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &database)

	funcs := template.FuncMap{"CreateFields": CreateFields, "InsertFields": InsertFields}
	tmpl, err := template.New("gen").Funcs(funcs).ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	tables := []Data{}
	for name, table := range database {
		data := Data{Name: name, Table: table}
		tables = append(tables, data)
	}

	templateData := TemplateData{
		Input:  databaseFile,
		Tables: tables}

	err = tmpl.ExecuteTemplate(os.Stdout, path.Base(templateFile), &templateData)
	if err != nil {
		log.Fatal(err)
	}
}
