package main

import (
	"testing"
)

func Test_isForeignKey(t *testing.T) {
	name := "catalog_series_id"
	if isForeignKey(name) != true {
		t.Errorf("Expected '%s' to be detected as foreign key", name)
	}

	name = "description"
	if isForeignKey(name) != false {
		t.Errorf("Expected '%s' to NOT be detected as foreign key", name)
	}
}

func Test_toObjName(t *testing.T) {
	tableName := "catalog_series_category"
	expect := "CatalogSeriesCategory"

	got := toObjName(tableName)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toUpdateFields(t *testing.T) {
	columns := []Column{
		Column{Name: "id", Type: "text"},
		Column{Name: "name", Type: "text"},
		Column{Name: "catalog_series_id", Type: "text"},
		Column{Name: "image_group", Type: "text"},
	}
	expect := `name=?,catalog_series_id=?,image_group=?`
	got := toUpdateFields(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toUpdateParams(t *testing.T) {
	columns := []Column{
		Column{Name: "id", Type: "text"},
		Column{Name: "name", Type: "text"},
		Column{Name: "catalog_series_id", Type: "text"},
		Column{Name: "image_group", Type: "text"},
	}
	expect := `input.Name,input.CatalogSeriesID,input.ImageGroup`
	got := toUpdateParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toSelectParams(t *testing.T) {
	columns := []Column{
		Column{Name: "id", Type: "text"},
		Column{Name: "name", Type: "text"},
		Column{Name: "catalog_series_id", Type: "text"},
		Column{Name: "image_group", Type: "text"},
	}
	expect := "id,name,catalog_series_id,image_group"
	got := toSelectParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toScanParams(t *testing.T) {
	columns := []Column{
		Column{Name: "id", Type: "text"},
		Column{Name: "name", Type: "text"},
		Column{Name: "catalog_series_id", Type: "text"},
		Column{Name: "image_group", Type: "text"},
	}
	expect := "&obj.ID,&obj.Name,&obj.CatalogSeriesID,&obj.ImageGroup"
	got := toScanParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}
