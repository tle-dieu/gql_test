package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tle-dieu/gql_test/internal/db/mysql"
	"github.com/tle-dieu/gql_test/internal/http/handler"
	"github.com/tle-dieu/gql_test/pkg/middleware/logger"
	mysqlMiddleware "github.com/tle-dieu/gql_test/pkg/middleware/mysql"
)

func main() {
	router := mux.NewRouter()
	mysqlClient := mysql.NewMySQLClient()
	serv := &http.Server{
		Handler:      mysqlMiddleware.WrapMysqlHandler(mysqlClient, logger.WrapHandlerWithLogging(router)),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/ads", handler.GetAds).Methods(http.MethodPost)
	// router.HandleFunc("/createAd", handler.CreateAd).Methods(http.MethodPost)
	// router.HandleFunc("/updateAd", handler.UpdateAd).Methods(http.MethodPost)
	// router.HandleFunc("/deleteAd", handler.DeleteAd).Methods(http.MethodPost)
	log.Println("Listening on :8080")
	log.Fatal(serv.ListenAndServe())
}
