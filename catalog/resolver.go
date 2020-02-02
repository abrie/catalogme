package catalog

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CatalogSeries(ctx context.Context, id string) (*CatalogSeries, error) {
	return &CatalogSeries{Name: "series1", Categories: []CatalogSeriesCategory{CatalogSeriesCategory{Name: "categorya"}}}, nil
}
func (r *queryResolver) CatalogSeriesCategory(ctx context.Context, id string) (*CatalogSeriesCategory, error) {
	return &CatalogSeriesCategory{Name: "category1"}, nil
}
func (r *queryResolver) CatalogSeriesCategoryPart(ctx context.Context, id string) (*CatalogSeriesCategoryPart, error) {
	return &CatalogSeriesCategoryPart{Name: "part1"}, nil
}
func (r *queryResolver) CatalogSeriesCategoryPartVersion(ctx context.Context, id string) (*CatalogSeriesCategoryPartVersion, error) {
	return &CatalogSeriesCategoryPartVersion{Name: "version1"}, nil
}
