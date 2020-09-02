package gqlgen

import "github.com/tle-dieu/gql_test/gqlgen/model"

func adInputToAd(input model.AdInput) *model.Ad {
	options := &model.Options{}
	if input.Options != nil {
		options.Bluetooth = input.Options.Bluetooth
		options.Gps = input.Options.Gps
	}
	return &model.Ad{
		Ref:     input.Ref,
		Brand:   input.Brand,
		Model:   input.Model,
		Price:   input.Price,
		Options: options,
	}

}
