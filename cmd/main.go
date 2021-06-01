package main

import (
	"fmt"
	"net/http"

	"github.com/dakaii/superduperpotato/controller"
	"github.com/dakaii/superduperpotato/internal"
	"github.com/dakaii/superduperpotato/service"
	"github.com/dakaii/superduperpotato/storage"
)

func main() {
	db := storage.InitDatabase()
	services := service.InitServices(db)
	controllers := controller.InitControllers(services)
	schema := controller.Schema(controllers)
	handler := controller.GraphqlHandlfunc(schema)

	http.Handle("/graphql", corsMiddleware(handler))
	fmt.Println("graphql api server is started at: http://localhost:" + internal.ServerPort + "/graphql")
	http.ListenAndServe(":"+internal.ServerPort, nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
