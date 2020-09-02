package graph

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/tle-dieu/gql_test/graph/models"
	"github.com/tle-dieu/gql_test/pkg/db/mysql"
)

type Resolver struct {
	Db *mysql.ClientMySQL
}

type AdResolver struct {
	Ad *models.Ad
}

type AdsResolver struct {
	Ad []models.Ad
}

type OptionsResolver struct {
	Options *models.Options
}

func (r *AdResolver) Ref(ctx context.Context) graphql.ID {
	return r.Ad.Ref
}

func (r *AdResolver) Brand(ctx context.Context) string {
	return r.Ad.Brand
}

func (r *AdResolver) Model(ctx context.Context) string {
	return r.Ad.Model
}

func (r *AdResolver) Price(ctx context.Context) int32 {
	return r.Ad.Price
}

func (r *AdResolver) Options(ctx context.Context) *OptionsResolver {
	return &OptionsResolver{
		Options: r.Ad.Options,
	}
}

func (r *OptionsResolver) Bluetooth(ctx context.Context) *bool {
	return r.Options.Bluetooth
}

func (r *OptionsResolver) Gps(ctx context.Context) *bool {
	return r.Options.Gps
}
