package client

import (
	"net/http"

	"github.com/gemalto/requester"
	"github.com/tle-dieu/ad_http_api/domain/model"
)

type Client struct {
	requester *requester.Requester
}

type CreateAdResponse struct {
	Ref string `json:"ref"`
}

func NewClient(apiURL string, httpClient *http.Client) *Client {
	return &Client{
		requester: requester.MustNew(
			requester.JSON(true),
			requester.URL(apiURL),
			requester.WithDoer(httpClient),
			requester.ExpectCode(http.StatusAccepted),
		),
	}
}

func (cli *Client) CreateAd(ad *model.Ad) (CreateAdResponse, error) {
	requesterOptions := []requester.Option{
		requester.Post("/createAd"),
		requester.Body(ad),
	}
	var resp CreateAdResponse
	r, _, err := cli.requester.Receive(
		&resp,
		requesterOptions...,
	)
	r.Body.Close()
	return resp, err
}
