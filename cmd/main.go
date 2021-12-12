package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/dakaii/superduperpotato/controller"
	"github.com/dakaii/superduperpotato/internal"
	"github.com/dakaii/superduperpotato/service"
	"github.com/dakaii/superduperpotato/storage"

	"github.com/dakaii/superduperpotato/controller/auth"
	"github.com/dakaii/superduperpotato/model"
)

func main() {
	db := storage.InitDatabase()
	services := service.InitServices(db)
	controllers := controller.InitControllers(services)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
    r.Use(corsMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	    
	r.Route("/api", func(r chi.Router) {
		r.Post("/signup", controllers.AuthController.Signup)
		r.Post("/login", controllers.AuthController.Login)
	})

	fmt.Println("server is now running at: http://localhost:" + internal.ServerPort)
	http.ListenAndServe(":"+internal.ServerPort, r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
