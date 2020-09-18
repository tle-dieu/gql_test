package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tle-dieu/gql_test/internal/db/mysql"
	model "github.com/tle-dieu/gql_test/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

//Fizzbuzz handle a fizzbuzz request (/fizzbuzz)
// func Fizzbuzz(w http.ResponseWriter, req *http.Request) {
// 	d, err := fizzbuzzGetDataRequest(req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := fizzbuzz.CheckData(d); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	contentType, err := fizzbuzzCheckResponseType(req.Header.Get("Accept-encoding"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotAcceptable)
// 		return
// 	}
// 	w.Header().Set("Content-type", contentType)
// 	w.Write([]byte(fizzbuzz.FizzbuzzAlgo(d)))
// }

func fizzbuzzCheckResponseType(acceptEncoding string) (string, error) {
	contentTypeResponse := [...]string{"application/model"}
	arrAcceptEncoding := strings.Split(acceptEncoding, ",")

	fmt.Println(acceptEncoding)
	if len(arrAcceptEncoding) == 0 || strings.TrimSpace(arrAcceptEncoding[0]) == "" {
		return contentTypeResponse[0], nil
	}
	for _, v := range arrAcceptEncoding {
		for _, contentType := range contentTypeResponse {
			if strings.TrimSpace(v) == contentType {
				return contentType, nil
			}
		}
	}
	return "", fmt.Errorf("Bad Accept-encoding, can be %v", contentTypeResponse)
}

func fizzbuzzGetAdsRequest(req *http.Request) (*model.Ad, error) {
	var err error

	if req.Body == nil || req.Body == http.NoBody {
		return nil, errors.New("Please send a request body")
	}
	if req.Header.Get("Content-Type") != "application/model" {
		return nil, errors.New("Content-Type must be application/model")
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("Unable to read message from request : " + err.Error())
	}
	adRequest := &model.Ad{}
	if err = proto.Unmarshal(body, adRequest); err != nil {
		return nil, err
	}
	return adRequest, nil
}

func GetAds(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ads, err := db.GetAllAds()
		if err != nil {
			panic(err) // @FIXME
		}
		response, err := proto.Marshal(ads)
		if err != nil {
			panic(err) // @FIXME
		}
		_, err = w.Write(response)
		if err != nil {
			log.Fatalf("Unable to write data into HTTP response : %v", err)
		}
	}
}
