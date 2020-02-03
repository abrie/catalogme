package catalog

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{ Datastore *Datastore }

func (r *Resolver) CatalogSeries() CatalogSeriesResolver {
	return &catalogSeriesResolver{r}
}
func (r *Resolver) CatalogSeriesCategory() CatalogSeriesCategoryResolver {
	return &catalogSeriesCategoryResolver{r}
}
func (r *Resolver) CatalogSeriesCategoryPart() CatalogSeriesCategoryPartResolver {
	return &catalogSeriesCategoryPartResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type catalogSeriesResolver struct{ *Resolver }

func (r *catalogSeriesResolver) CatalogSeriesCategoryList(ctx context.Context, obj *CatalogSeries) ([]*CatalogSeriesCategory, error) {
	return r.Datastore.List_CatalogSeriesCategory(*obj.ID)
}

type catalogSeriesCategoryResolver struct{ *Resolver }

func (r *catalogSeriesCategoryResolver) CatalogSeriesCategoryPartList(ctx context.Context, obj *CatalogSeriesCategory) ([]*CatalogSeriesCategoryPart, error) {
	return r.Datastore.List_CatalogSeriesCategoryPart(*obj.ID)
}

type catalogSeriesCategoryPartResolver struct{ *Resolver }

func (r *catalogSeriesCategoryPartResolver) CatalogSeriesCategoryPartVersionList(ctx context.Context, obj *CatalogSeriesCategoryPart) ([]*CatalogSeriesCategoryPartVersion, error) {
	return r.Datastore.List_CatalogSeriesCategoryPartVersion(*obj.ID)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CatalogSeries(ctx context.Context, id string) (*CatalogSeries, error) {
	return r.Datastore.GetCatalogSeries(id)
}
