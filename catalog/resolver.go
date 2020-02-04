package catalog

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Datastore *Datastore
}

func (r *Resolver) CatalogSeries() CatalogSeriesResolver {
	return &catalogSeriesResolver{r}
}
func (r *Resolver) CatalogSeriesCategory() CatalogSeriesCategoryResolver {
	return &catalogSeriesCategoryResolver{r}
}
func (r *Resolver) CatalogSeriesCategoryPart() CatalogSeriesCategoryPartResolver {
	return &catalogSeriesCategoryPartResolver{r}
}
func (r *Resolver) PortfolioGroup() PortfolioGroupResolver {
	return &portfolioGroupResolver{r}
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

type portfolioGroupResolver struct{ *Resolver }

func (r *portfolioGroupResolver) PortfolioGroupItemList(ctx context.Context, obj *PortfolioGroup) ([]*PortfolioGroupItem, error) {
	return r.Datastore.List_PortfolioGroupItem(*obj.ID)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CatalogSeries(ctx context.Context, id string) (*CatalogSeries, error) {
	return r.Datastore.GetCatalogSeries(id)
}
func (r *queryResolver) PortfolioGroup(ctx context.Context, id string) (*PortfolioGroup, error) {
	return r.Datastore.GetPortfolioGroup(id)
}
func (r *queryResolver) ServiceItem(ctx context.Context, id string) (*ServiceItem, error) {
	return r.Datastore.GetServiceItem(id)
}
func (r *queryResolver) RestorationItem(ctx context.Context, id string) (*RestorationItem, error) {
	return r.Datastore.GetRestorationItem(id)
}
func (r *queryResolver) AboutTopic(ctx context.Context, id string) (*AboutTopic, error) {
	return r.Datastore.GetAboutTopic(id)
}
