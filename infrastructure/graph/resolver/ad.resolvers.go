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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
