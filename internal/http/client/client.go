package ads

import (
	"net/http"

	"github.com/gemalto/requester"
	"github.com/golang/protobuf/proto"
	"github.com/tle-dieu/gql_test/pkg/protobuf"
)

type Client struct {
	requester *requester.Requester
}

func NewClient(apiUrl string, httpClient *http.Client) *Client {
	return &Client{
		requester: requester.MustNew(
			requester.URL(apiUrl),
			requester.WithDoer(httpClient),
			requester.ExpectCode(http.StatusOK),
			requester.AddHeader("Connection", "keep-alive"),
			requester.Accept("application/protobuf"),
			requester.ContentType("application/protobuf"),
			requester.WithMarshaler(proto.Marshaler{}),
		),
	}
}

func (cli *Client) CreateAd(ad *protobuf.Ad) error {
	createAdRequest := protobuf.CreateAdRequest{Ad: ad}
	requesterOptions := []requester.Option{
		requester.Post("/createAd"),
		requester.Body(createAdRequest),
	}
	_, _, err := cli.requester.Receive(
		requester.ExpectCode(http.StatusAccepted),
		requesterOptions...,
	)
	return err
}
