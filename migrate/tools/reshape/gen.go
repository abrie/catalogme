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
type Columns []Column

type Table struct {
	Name    string
	Columns Columns
}

type TemplateData struct {
	InputDatabase string
	Tables        []Table
}

type Database map[string]Columns

func shouldDropColumn(table Table, column Column) bool {
	if table.Name == "catalog_series_category_part" {
		if column["name"] == "tag" {
			return true
		}
	}

	return false
}

func FieldsForCreate(table Table) (string, error) {
	var list []string
	for _, column := range table.Columns {
		if shouldDropColumn(table, column) {
			continue
		}

		/* Modification: If price column, set data type to integer */
		var datatype string
		if column["name"] == "price" {
			datatype = "integer"
		} else {
			datatype = column["type"]
		}
		/***/

		list = append(list, fmt.Sprintf(`"%s" %s`, column["name"], datatype))
	}
	return strings.Join(list, ", "), nil
}

func FieldsForInsert(table Table) (string, error) {
	var list []string
	for _, column := range table.Columns {
		if shouldDropColumn(table, column) {
			continue
		}
		list = append(list, fmt.Sprintf(`"%s"`, column["name"]))
	}
	return strings.Join(list, ", "), nil
}

func FieldsForSelect(table Table) (string, error) {
	var list []string
	for _, column := range table.Columns {
		if shouldDropColumn(table, column) {
			continue
		}

		/* Modification: If price column, multiply by 100 to get cent denominated price */
		var name string
		if column["name"] == "price" {
			name = fmt.Sprintf(`"%s"*100`, column["name"])
		} else {
			name = fmt.Sprintf(`"%s"`, column["name"])
		}
		/***/

		list = append(list, name)
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

	funcs := template.FuncMap{"FieldsForCreate": FieldsForCreate, "FieldsForSelect": FieldsForSelect, "FieldsForInsert": FieldsForInsert}
	tmpl, err := template.New("gen").Funcs(funcs).ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	tables := []Table{}
	for name, columns := range database {

		/* Modification: Filter out unused tables */
		if name == "home" {
			continue
		}
		if name == "home_item" {
			continue
		}
		/* */

		table := Table{Name: name, Columns: columns}
		tables = append(tables, table)
	}

	templateData := TemplateData{
		InputDatabase: databaseFile,
		Tables:        tables}

	err = tmpl.ExecuteTemplate(os.Stdout, path.Base(templateFile), &templateData)
	if err != nil {
		log.Fatal(err)
	}
}
