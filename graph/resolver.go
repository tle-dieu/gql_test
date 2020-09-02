package graph

import (
	"context"

	"github.com/tle-dieu/gql_test/graph/model"
	"github.com/tle-dieu/gql_test/pkg/db/mysql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type AdResolver struct {
	Db *mysql.ClientMySQL
}

func (r *AdResolver) CreateAd(ctx context.Context, args struct{ model.AdInput }) (*model.Ad, error) {
	if err := r.Db.SaveAd(args.AdInput); err != nil {
		return nil, gqlerror.Errorf("error while creating Ad: " + err.Error())
	}
	return adInputToAd(args.AdInput), nil
}

func (r *AdResolver) UpdateAd(ctx context.Context, args struct{ input model.AdInput }) (*model.Ad, error) {
	if err := r.Db.UpdateAd(args.input); err != nil {
		return nil, gqlerror.Errorf("error while updating Ad: " + err.Error())
	}
	return adInputToAd(args.input), nil
}

func (r *AdResolver) DeleteAd(ctx context.Context, args struct{ ref string }) (bool, error) {
	if err := r.Db.DeleteAd(args.ref); err != nil {
		return false, gqlerror.Errorf("error while deleting Ad: " + err.Error())
	}
	return true, nil
}

func (r *AdResolver) Ads(ctx context.Context) ([]*model.Ad, error) {
	ads, err := r.Db.GetAllAds()
	if err != nil {
		return nil, gqlerror.Errorf("error while getting Ads: " + err.Error())
	}
	return ads, nil
}
