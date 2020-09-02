package graph

import (
	"context"

	"github.com/vektah/gqlparser/gqlerror"
)

func (r *Resolver) Ads(ctx context.Context) ([]*AdResolver, error) {
	ads, err := r.Db.GetAllAds()
	if err != nil {
		return nil, gqlerror.Errorf("error while getting Ads: " + err.Error())
	}
	var adsResolver []*AdResolver
	for _, ad := range ads {
		adsResolver = append(adsResolver, &AdResolver{&ad})
	}
	return adsResolver, nil
}
