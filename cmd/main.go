package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	core "github.com/Mboukhal/FactoryBase/core"
	sqlc "github.com/Mboukhal/FactoryBase/internal/db"
	"github.com/Mboukhal/FactoryBase/internal/settings"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// ctx := context.Background()

	app_env := os.Getenv("APP_ENV")
	if app_env == "" {
		app_env = os.Getenv("NODE_ENV")
	}
	if app_env == "" {
		app_env = "production"
	}

	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := settings.OpenDB()
	if err != nil {
		panic("Failed to connect to the database:" + err.Error())
	}
	defer db.Close()

	// Create queries instance
	queries := sqlc.New(db)

	// // insert a test user
	// _, err = queries.CreateProfile(ctx, CreateProfileParams{
	// 	ID:       uuid.NewString(),
	// 	Username: sql.NullString{String: "testuser", Valid: true},
	// 	Phone:    sql.NullString{String: "1234567890", Valid: true},
	// 	Email:    "testuser@example.com",
	// 	Role:     "learner",
	// })

	// if err != nil {
	// 	log.Println("Error inserting test user:", err)
	// }

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	// r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	// Inject queries into context
	r.Use(settings.WithQueries(queries))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chimiddleware.Timeout(60 * time.Second))

	isProduction := app_env == "production"
	if isProduction {
		settings.ProductionSettings(r)
	} else {
		settings.DevelopmentSettings(r)
	}

	core.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s in %s mode", port, app_env)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
