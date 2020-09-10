package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tle-dieu/gql_test/internal/http/handler"
	"github.com/tle-dieu/gql_test/pkg/http/logger"
)

func main() {
	router := mux.NewRouter()
	serv := &http.Server{
		Handler:      logger.WrapHandlerWithLogging(router),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/getAds", handler.GetAds).Methods(http.MethodPost)
	// router.HandleFunc("/createAd", handler.CreateAd).Methods(http.MethodPost)
	// router.HandleFunc("/updateAd", handler.UpdateAd).Methods(http.MethodPost)
	// router.HandleFunc("/deleteAd", handler.DeleteAd).Methods(http.MethodPost)
	log.Println("Listening on :8080")
	log.Fatal(serv.ListenAndServe())
}
