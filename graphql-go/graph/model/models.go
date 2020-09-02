package model

type Ad struct {
	Ref     string   `json:"ref"`
	Brand   string   `json:"brand"`
	Model   string   `json:"model"`
	Price   int      `json:"price"`
	Options *Options `json:"options"`
}

type AdInput struct {
	Ref     string        `json:"ref"`
	Brand   string        `json:"brand"`
	Model   string        `json:"model"`
	Price   int           `json:"price"`
	Options *OptionsInput `json:"options"`
}

type Options struct {
	Bluetooth *bool `json:"bluetooth"`
	Gps       *bool `json:"gps"`
}

type OptionsInput struct {
	Bluetooth *bool `json:"bluetooth"`
	Gps       *bool `json:"gps"`
}
