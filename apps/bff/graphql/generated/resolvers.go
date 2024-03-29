package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"github.com/v3nooonn/trytry/apps/bff/graphql/model"
)

// Brands is the resolver for the brands field.
func (r *queryResolver) Brands(ctx context.Context, id *string) ([]*model.Brand, error) {
	panic(fmt.Errorf("not implemented: Brands - brands"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
