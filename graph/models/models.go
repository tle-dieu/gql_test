package models

import (
	"github.com/graph-gophers/graphql-go"
)

type Ad struct {
	Ref     graphql.ID
	Brand   string
	Model   string
	Price   int32
	Options *Options
}

type Options struct {
	Bluetooth *bool
	Gps       *bool
}

type AdInput struct {
	Ref     graphql.ID
	Brand   string
	Model   string
	Price   int32
	Options *OptionsInput
}

type OptionsInput struct {
	Bluetooth *bool
	Gps       *bool
}
