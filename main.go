package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rssagg/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfiguration struct {
	Database *database.Queries
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DB url not found in the env")
	}

	connection, error := sql.Open("postgres", dbURL)
	if error != nil{
		log.Printf("Failed to connect to database: %v", error)
	}

	queries:= database.New(connection)

	apiCpnfiguration := ApiConfiguration{
		Database: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	  v1Router := chi.NewRouter()
	  v1Router.Get("/health", handlerReadiness)
	  v1Router.Get("/error", handlerError)
	  v1Router.Post("/user", apiCpnfiguration.handlerCreateUser)

	  router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	fmt.Printf("Server started at %v", portString)
	err := srv.ListenAndServe()

	if err != nil{
		fmt.Println("Connection error:", err)
	}
}