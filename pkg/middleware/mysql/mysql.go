package mysql

import (
	"fmt"
	"net/http"

	"github.com/gol4ng/httpware/v3"
	"github.com/tle-dieu/gql_test/internal/db/mysql"
)

type ResponseWriterClient struct {
	http.ResponseWriter
	Db *mysql.Client
}

func WrapMysqlHandler(cli *mysql.Client) httpware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("yes")
			responseWriter := &ResponseWriterClient{
				ResponseWriter: w,
				Db:             cli,
			}
			next.ServeHTTP(responseWriter, r)
		})
	}
}
