package catalog

type CatalogSeries struct {
	ID          string
	Name        string
	Description string
	Categories  []CatalogSeriesCategory
}

type CatalogSeriesCategory struct {
	ID          string
	Name        string
	Description string
	Parts       []CatalogSeriesCategoryPart
}

type CatalogSeriesCategoryPart struct {
	ID          string
	Name        string
	Description string
	Versions    []CatalogSeriesCategoryPartVersion
}

type CatalogSeriesCategoryPartVersion struct {
	ID          string
	Name        string
	Description string
	Price       int
}
