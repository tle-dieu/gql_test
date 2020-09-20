package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gol4ng/httpware/v3"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger-http/middleware"
	"github.com/gol4ng/logger/formatter"
	logger_handler "github.com/gol4ng/logger/handler"
	"github.com/gorilla/mux"
	"github.com/tle-dieu/gql_test/internal/db/mysql"
	"github.com/tle-dieu/gql_test/internal/http/handler"
)

func main() {
	router := mux.NewRouter()
	l := logger.NewLogger(logger_handler.Stream(os.Stdout, formatter.NewDefaultFormatter(formatter.WithContext(true))))
	mysqlClient, err := mysql.NewClient("mysql", "localhost", 3307, "root", "password", "local-db")
	if err != nil {
		l.Error("error while connecting to mysql", logger.Error("error", err))
		return
	}
	stack := httpware.MiddlewareStack(
		middleware.CorrelationId(),
		middleware.Logger(l),
	)
	err = mysqlClient.Migrate()
	if err != nil {
		l.Error("error while migrating mysql", logger.Error("error", err))
		return
	}

	router.HandleFunc("/ads", handler.GetAds(*mysqlClient)).Methods(http.MethodGet)
	router.HandleFunc("/createAd", handler.CreateAd(*mysqlClient)).Methods(http.MethodPost)
	router.HandleFunc("/updateAd", handler.UpdateAd(*mysqlClient)).Methods(http.MethodPut)
	router.HandleFunc("/deleteAd/{ad_ref}", handler.DeleteAd(*mysqlClient)).Methods(http.MethodDelete)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack.DecorateHandler(router),
		BaseContext: func(listener net.Listener) context.Context {
			return logger.InjectInContext(context.Background(), l)
		},
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// router.HandleFunc("/createAd", handler.CreateAd).Methods(http.MethodPost)
	// router.HandleFunc("/updateAd", handler.UpdateAd).Methods(http.MethodPost)
	// router.HandleFunc("/deleteAd", handler.DeleteAd).Methods(http.MethodPost)
	l.Info("Listening on :8081")
	l.Error(server.ListenAndServe().Error())
}
