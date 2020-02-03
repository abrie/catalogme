package catalog

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Datastore struct {
	DB *sql.DB
}

func (datastore *Datastore) Ping() {
	err := datastore.DB.Ping()
	if err != nil {
		log.Println("Datastore failure.", err)
	}
}

func OpenDatastore(dbFile string) *Datastore {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	return &Datastore{DB: db}
}

func (datastore *Datastore) Close() {
	datastore.DB.Close()
}

func (datastore *Datastore) GetCatalogSeriesCategories(rowid string) ([]*CatalogSeriesCategory, error) {
	var (
		catalog_series_categories []*CatalogSeriesCategory
	)
	rows, err := datastore.DB.Query("select ROWID, name, description, shortname, image_group from catalog_series_category where catalog_series_id=?", rowid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		catalog_series_category := CatalogSeriesCategory{}
		err := rows.Scan(&catalog_series_category.ID, &catalog_series_category.Name, &catalog_series_category.Description, &catalog_series_category.Shortname, &catalog_series_category.ImageGroup)
		if err != nil {
			log.Fatal(err)
		}
		catalog_series_categories = append(catalog_series_categories, &catalog_series_category)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return catalog_series_categories, nil
}

func (datastore *Datastore) GetCatalogSeries(rowid string) (*CatalogSeries, error) {
	var (
		catalog_series CatalogSeries
	)

	err := datastore.DB.QueryRow("select ROWID, name, description, shortname, image_group from catalog_series where ROWID=?", rowid).Scan(&catalog_series.ID, &catalog_series.Name, &catalog_series.Description, &catalog_series.Shortname, &catalog_series.ImageGroup)
	if err != nil {
		log.Fatal(err)
	}

	return &catalog_series, nil
}
