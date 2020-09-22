package transformer

import "github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated/model"

func AdInputToAd(input model.AdInput, ref string) *model.Ad {
	options := &model.Options{}
	if input.Options != nil {
		options.Bluetooth = input.Options.Bluetooth
		options.Gps = input.Options.Gps
	}
	return &model.Ad{
		Ref:     ref,
		Brand:   input.Brand,
		Model:   input.Model,
		Price:   input.Price,
		Options: options,
	}
}
