package client

import (
	"net/http"

	"github.com/gemalto/requester"
	"github.com/tle-dieu/ad_graphql_api/infrastructure/graph/generated/model"
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
			requester.ExpectCode(http.StatusOK),
			// should be this instead of StatusOK
			// requester.ExpectCode(http.StatusAccepted),
		),
	}
}

func (cli *Client) CreateAd(ad *model.Ad) (CreateAdResponse, error) {
	requesterOptions := []requester.Option{
		requester.Post("/createAd"),
		requester.Body(ad),
	}
	var adResponse CreateAdResponse
	resp, _, err := cli.requester.Receive(
		&adResponse,
		requesterOptions...,
	)
	if err != nil {
		return adResponse, err
	}

	resp.Body.Close()
	return adResponse, nil
}
