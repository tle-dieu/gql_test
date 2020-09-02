package graph

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/tle-dieu/gql_test/graph/models"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *Resolver) DeleteAd(ctx context.Context, args struct{ Ref graphql.ID }) (bool, error) {
	if err := r.Db.DeleteAd(string(args.Ref)); err != nil {
		return false, gqlerror.Errorf("error while deleting Ad: " + err.Error())
	}
	return true, nil
}

func (r *Resolver) CreateAd(ctx context.Context, args struct{ Ad models.AdInput }) (*AdResolver, error) {
	if err := r.Db.SaveAd(args.Ad); err != nil {
		return nil, gqlerror.Errorf("error while creating Ad: " + err.Error())
	}
	return &AdResolver{adInputToAd(args.Ad)}, nil
}

func (r *Resolver) UpdateAd(ctx context.Context, args struct{ Ad models.AdInput }) (*AdResolver, error) {
	if err := r.Db.UpdateAd(args.Ad); err != nil {
		return nil, gqlerror.Errorf("error while updating Ad: " + err.Error())
	}
	return &AdResolver{adInputToAd(args.Ad)}, nil
}
