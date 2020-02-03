package main

import (
	"testing"
)

func Test_toObjName(t *testing.T) {
	tableName := "catalog_series_category"
	expect := "CatalogSeriesCategory"

	got := toObjName(tableName)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toSelectParams(t *testing.T) {
	columns := []string{"id", "name", "description", "image_group"}
	expect := "ROWID,name,description,image_group"
	got := toSelectParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toScanParams(t *testing.T) {
	columns := []string{"id", "name", "description", "image_group"}
	expect := "&obj.ID,&obj.Name,&obj.Description,&obj.ImageGroup"
	got := toScanParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}
