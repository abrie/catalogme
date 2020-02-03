package catalog

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) CatalogSeries() CatalogSeriesResolver {
	return &catalogSeriesResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type catalogSeriesResolver struct{ *Resolver }

func (r *catalogSeriesResolver) CatalogSeriesCategoryList(ctx context.Context, obj *CatalogSeries) ([]*CatalogSeriesCategory, error) {
	return r.Datastore.GetCatalogSeriesCategoryList(*obj.ID)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CatalogSeries(ctx context.Context, id string) (*CatalogSeries, error) {
	return r.Datastore.GetCatalogSeries(id)
}
