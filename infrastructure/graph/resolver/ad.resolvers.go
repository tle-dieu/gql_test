package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated/model"
)

func (r *mutationResolver) CreateAd(ctx context.Context, input model.AdInput) (*model.Ad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAd(ctx context.Context, input model.AdInput) (*model.Ad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAd(ctx context.Context, ref string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ads(ctx context.Context) ([]model.Ad, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
