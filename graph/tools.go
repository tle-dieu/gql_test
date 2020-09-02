package graph

import "github.com/tle-dieu/gql_test/graph/models"

func adInputToAd(input models.AdInput) *models.Ad {
	options := &models.Options{
		Bluetooth: new(bool),
		Gps:       new(bool),
	}
	if input.Options != nil {
		if input.Options.Bluetooth != nil {
			options.Bluetooth = input.Options.Bluetooth
		}
		if input.Options.Gps != nil {
			options.Bluetooth = input.Options.Gps
		}
	}
	return &models.Ad{
		Ref:     input.Ref,
		Brand:   input.Brand,
		Model:   input.Model,
		Price:   input.Price,
		Options: options,
	}
}
