package handler

import (
	"net/http"

	"github.com/tle-dieu/gql_test/internal/db/mysql"
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

func GetAds(db mysql.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ads, err := db.GetAllAds()
		if err != nil {
			// return nil, gqlerror.Errorf("error while getting Ads: " + err.Error())
			panic(err)
		}
		// return ads, nil
	}
}
