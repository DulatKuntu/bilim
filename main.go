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
	"github.com/rs/cors"
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
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	srv := &http.Server{
		Handler:      corsWrapper.Handler(r),
		Addr:         ":" + config.Port,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	go srv.ListenAndServe()
	log.Fatal(srvStatis.ListenAndServe())
}

func InitStaticRoutes(staticRouter *mux.Router, handler *handler.AppHandler) {
	//staticRouter.HandleFunc("/zip/{fileName}", handler.ZIPHandler).Methods("POST")
	staticRouter.HandleFunc("/{locationType}/{fileName}", handler.RouterImageHandler).Methods("GET")
}

func InitRoutes(r *mux.Router, handler *handler.AppHandler) {
	unauthed := r.PathPrefix("/unauthed").Subrouter()
	buddy := r.PathPrefix("/buddy").Subrouter()
	authed := r.PathPrefix("/authed").Subrouter()
	unauthed.HandleFunc("/signup", handler.SignUp).Methods("POST")
	unauthed.HandleFunc("/signupMentor", handler.SignUpMentor).Methods("POST")
	unauthed.HandleFunc("/signin", handler.SignIn).Methods("POST")
	unauthed.HandleFunc("/signinMentor", handler.SignInMentor).Methods("POST")
	authed.HandleFunc("/getProfile", handler.GetProfile).Methods("GET")
	authed.HandleFunc("/getMentors", handler.GetMentors).Methods("GET")
	authed.HandleFunc("/getMentor", handler.GetProfileMentor).Methods("GET")
	authed.HandleFunc("/updateProfile", handler.UpdateProfile).Methods("POST")
	buddy.HandleFunc("/postBuddy", handler.PostBuddy).Methods("POST")
	authed.HandleFunc("/addInterest", handler.AddInterests).Methods("POST")
	authed.HandleFunc("/addMentorInterest", handler.AddMentorInterests).Methods("POST")
	authed.HandleFunc("/getPosts", handler.GetPosts).Methods("GET")
}
