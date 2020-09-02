package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tle-dieu/gql_test/gqlgen/graph/generated"
	"github.com/tle-dieu/gql_test/gqlgen/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateAd(ctx context.Context, input model.AdInput) (*model.Ad, error) {
	if err := r.Db.SaveAd(input); err != nil {
		return nil, gqlerror.Errorf("error while creating Ad:", err)
	}
	return AdInputToAd(input), nil
}

func (r *mutationResolver) UpdateAd(ctx context.Context, input model.AdInput) (*model.Ad, error) {
	if err := r.Db.UpdateAd(input); err != nil {
		return nil, gqlerror.Errorf("error while updating Ad:", err)
	}
	return AdInputToAd(input), nil
}

func (r *mutationResolver) DeleteAd(ctx context.Context, ref string) (bool, error) {
	if err := r.Db.DeleteAd(ref); err != nil {
		return false, gqlerror.Errorf("error while deleting Ad:", err)
	}
	return true, nil
}

func (r *queryResolver) Ads(ctx context.Context) ([]*model.Ad, error) {
	ads, err := r.Db.GetAllAds()
	if err != nil {
		return nil, gqlerror.Errorf("error while getting Ads:", err)
	}
	return ads, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
