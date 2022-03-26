package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DulatKuntu/bilim/controller"
	"github.com/DulatKuntu/bilim/database"
	"github.com/DulatKuntu/bilim/handler"
	"github.com/DulatKuntu/bilim/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config, err := utils.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Print("StartUp Dev: version 0.0.10")

	client, err := database.GetDB(config.Dburi)
	if err != nil {
		log.Printf("Mongo db connection err %v", err)
		return
	}

	defer client.Disconnect(context.TODO())
	database := client.Database(config.Dbname)
	repo := controller.NewDatabaseRepository(database, client)
	handler := handler.NewHandler(repo)

	r := mux.NewRouter()
	InitRoutes(r, handler)
	TestRoutes(r, handler)

	// Static Routers
	// # Used to handle static files such as: images, videos
	staticRouter := mux.NewRouter()
	InitStaticRoutes(staticRouter, handler)

	srvStatis := &http.Server{
		Handler:      staticRouter,
		Addr:         ":" + config.StaticPort,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + config.Port,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	go srv.ListenAndServe()
	log.Fatal(srvStatis.ListenAndServe())
}

func InitStaticRoutes(staticRouter *mux.Router, handler *handler.AppHandler) {
	//staticRouter.HandleFunc("/zip/{fileName}", handler.ZIPHandler).Methods("POST")
	//staticRouter.HandleFunc("/{locationType}/{imageType}/{fileName}", handler.RouterImageHandler).Methods("GET")
}

func InitRoutes(r *mux.Router, handler *handler.AppHandler) {
	//unauthed := r.PathPrefix("/unauthed").Subrouter()
	//authed := r.PathPrefix("/authed").Subrouter()

	//unauthed.HandleFunc("/signin", handler.SingIn).Methods("POST")

}

func TestRoutes(r *mux.Router, handler *handler.AppHandler) {

}
