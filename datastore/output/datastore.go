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

func (datastore *Datastore) GetCatalogSeriesCategoryPartVersion(id string) (*CatalogSeriesCategoryPartVersion, error) {
	var (
		obj CatalogSeriesCategoryPartVersion
	)

	err := datastore.DB.QueryRow("select catalog_series_category_part_id,code,description,image_group,internalcode,price,ROWID from catalog_series_category_part_version where ROWID=?", id).Scan(&obj.CatalogSeriesCategoryPartID, &obj.Code, &obj.Description, &obj.ImageGroup, &obj.Internalcode, &obj.Price, &obj.ID)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) List_CatalogSeriesCategoryPartVersion(id string) ([]*CatalogSeriesCategoryPartVersion, error) {
	var (
		list []*CatalogSeriesCategoryPartVersion
	)
	rows, err := datastore.DB.Query("select catalog_series_category_part_id,code,description,image_group,internalcode,price,ROWID from catalog_series_category_part_version where catalog_series_category_part_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj := CatalogSeriesCategoryPartVersion{}
		err := rows.Scan(&obj.CatalogSeriesCategoryPartID, &obj.Code, &obj.Description, &obj.ImageGroup, &obj.Internalcode, &obj.Price, &obj.ID)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, &obj)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return list, nil
}

func (datastore *Datastore) GetMarketCarforsale(id string) (*MarketCarforsale, error) {
	var (
		obj MarketCarforsale
	)

	err := datastore.DB.QueryRow("select description,image_group,name,ROWID,shortname from market_carforsale where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetPortfolioGroup(id string) (*PortfolioGroup, error) {
	var (
		obj PortfolioGroup
	)

	err := datastore.DB.QueryRow("select image_group,name,ROWID,shortname from portfolio_group where ROWID=?", id).Scan(&obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetPortfolioGroupItem(id string) (*PortfolioGroupItem, error) {
	var (
		obj PortfolioGroupItem
	)

	err := datastore.DB.QueryRow("select description,image_group,name,portfolio_group_id,ROWID,shortname from portfolio_group_item where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.PortfolioGroupID, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) List_PortfolioGroupItem(id string) ([]*PortfolioGroupItem, error) {
	var (
		list []*PortfolioGroupItem
	)
	rows, err := datastore.DB.Query("select description,image_group,name,portfolio_group_id,ROWID,shortname from portfolio_group_item where portfolio_group_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj := PortfolioGroupItem{}
		err := rows.Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.PortfolioGroupID, &obj.ID, &obj.Shortname)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, &obj)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return list, nil
}

func (datastore *Datastore) GetRestorationItem(id string) (*RestorationItem, error) {
	var (
		obj RestorationItem
	)

	err := datastore.DB.QueryRow("select description,image_group,name,ROWID,shortname from restoration_item where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetServiceItem(id string) (*ServiceItem, error) {
	var (
		obj ServiceItem
	)

	err := datastore.DB.QueryRow("select description,image_group,name,ROWID,shortname from service_item where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetAboutTopic(id string) (*AboutTopic, error) {
	var (
		obj AboutTopic
	)

	err := datastore.DB.QueryRow("select description,image_group,name,ROWID,shortname from about_topic where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetCatalogSeriesCategory(id string) (*CatalogSeriesCategory, error) {
	var (
		obj CatalogSeriesCategory
	)

	err := datastore.DB.QueryRow("select catalog_series_id,description,image_group,name,ROWID,shortname from catalog_series_category where ROWID=?", id).Scan(&obj.CatalogSeriesID, &obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) List_CatalogSeriesCategory(id string) ([]*CatalogSeriesCategory, error) {
	var (
		list []*CatalogSeriesCategory
	)
	rows, err := datastore.DB.Query("select catalog_series_id,description,image_group,name,ROWID,shortname from catalog_series_category where catalog_series_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj := CatalogSeriesCategory{}
		err := rows.Scan(&obj.CatalogSeriesID, &obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, &obj)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return list, nil
}

func (datastore *Datastore) GetCatalogSeriesCategoryPart(id string) (*CatalogSeriesCategoryPart, error) {
	var (
		obj CatalogSeriesCategoryPart
	)

	err := datastore.DB.QueryRow("select catalog_series_category_id,code,description,image_group,internalcode,price,ROWID,tag from catalog_series_category_part where ROWID=?", id).Scan(&obj.CatalogSeriesCategoryID, &obj.Code, &obj.Description, &obj.ImageGroup, &obj.Internalcode, &obj.Price, &obj.ID, &obj.Tag)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) List_CatalogSeriesCategoryPart(id string) ([]*CatalogSeriesCategoryPart, error) {
	var (
		list []*CatalogSeriesCategoryPart
	)
	rows, err := datastore.DB.Query("select catalog_series_category_id,code,description,image_group,internalcode,price,ROWID,tag from catalog_series_category_part where catalog_series_category_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj := CatalogSeriesCategoryPart{}
		err := rows.Scan(&obj.CatalogSeriesCategoryID, &obj.Code, &obj.Description, &obj.ImageGroup, &obj.Internalcode, &obj.Price, &obj.ID, &obj.Tag)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, &obj)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return list, nil
}

func (datastore *Datastore) GetImage(id string) (*Image, error) {
	var (
		obj Image
	)

	err := datastore.DB.QueryRow("select anchor,group,href,large_src,ROWID,sequence,small_src,tag from image where ROWID=?", id).Scan(&obj.Anchor, &obj.Group, &obj.Href, &obj.LargeSrc, &obj.ID, &obj.Sequence, &obj.SmallSrc, &obj.Tag)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}

func (datastore *Datastore) GetCatalogSeries(id string) (*CatalogSeries, error) {
	var (
		obj CatalogSeries
	)

	err := datastore.DB.QueryRow("select description,image_group,name,ROWID,shortname from catalog_series where ROWID=?", id).Scan(&obj.Description, &obj.ImageGroup, &obj.Name, &obj.ID, &obj.Shortname)
	if err != nil {
		log.Fatal(err)
	}

	return &obj, nil
}
