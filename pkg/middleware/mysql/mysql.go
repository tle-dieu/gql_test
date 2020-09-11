package mysql

import (
	"github.com/tle-dieu/gql_test/internal/db/mysql"
	"net/http"
)

type responseWriterClient struct {
	http.ResponseWriter
	Cli *mysql.ClientMySQL
}

func WrapMysqlHandler(cli *mysql.ClientMySQL, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseWriter := &responseWriterClient{
			ResponseWriter: w,
			Cli:            cli,
		}
		handler.ServeHTTP(responseWriter, r)
	})
}
