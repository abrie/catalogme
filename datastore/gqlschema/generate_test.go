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

func Test_toSelectParams(t *testing.T) {
	columns := []string{"id", "name", "catalog_series_id", "description", "image_group"}
	expect := "ROWID,name,catalog_series_id,description,image_group"
	got := toSelectParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}

func Test_toScanParams(t *testing.T) {
	columns := []string{"id", "name", "catalog_series_id", "description", "image_group"}
	expect := "&obj.ID,&obj.Name,&obj.CatalogSeriesID,&obj.Description,&obj.ImageGroup"
	got := toScanParams(columns)
	if got != expect {
		t.Errorf("Expected '%s', Got '%s'", expect, got)
	}
}
