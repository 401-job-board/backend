// @title Fan fit feed service
// @version 0.1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @description ## Feed
// @description Gives us details about workouts
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @BasePath /
package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	// API Routes
	"github.com/fanfit/feed/api/handlers/healthcheck"
	postHandlers "github.com/fanfit/feed/api/handlers/posts"
	"github.com/fanfit/feed/api/middleware"

	// Tags
	// Workout Tag
	postRepository "github.com/fanfit/feed/models/posts/repository"
	postServicePackage "github.com/fanfit/feed/models/posts/service"

	"github.com/fanfit/feed/server"
	"github.com/gin-gonic/gin"
)

type envVars struct {
	dbUserName string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	dbSchema   string
}

func main() {
	envVars, err := loadEnvVars()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while loading Env variables: %s", err.Error()))
	}
	dbURL := prepareDbURL(envVars)

	// Instantiate service for each tag
	workoutStore, err := postRepository.NewWorkoutStore(dbURL)
	if err != nil {
		log.Fatal(fmt.Printf("Error while creating userStore: %s", err.Error()))
	}
	workoutService := postServicePackage.New(workoutStore)

	// Initialize the middleware and routes
	engine := gin.Default()
	router := server.GenerateRouter(engine)

	// Set routes for each tag

	healthcheck.Routes(router)
	router.Use(middleware.VerifyToken)
	postHandlers.Routes(router, workoutService)

	server.Orchestrate(engine, workoutStore)
}

func loadEnvVars() (*envVars, error) {
	dbUsername, envPresent := os.LookupEnv("DB_USERNAME")
	if !envPresent {
		return nil, errors.New("DB_USERNAME environment variable missing")
	}

	dbPassword, envPresent := os.LookupEnv("DB_PASSWORD")
	if !envPresent {
		return nil, errors.New("DB_PASSWORD environment variable missing")
	}

	dbHost, envPresent := os.LookupEnv("DB_HOST")
	if !envPresent {
		return nil, errors.New("DB_HOST environment variable missing")
	}

	dbPort, envPresent := os.LookupEnv("DB_PORT")
	if !envPresent {
		return nil, errors.New("DB_PORT environment variable missing")
	}

	dbName, envPresent := os.LookupEnv("DB_NAME")
	if !envPresent {
		return nil, errors.New("DB_NAME environment variable missing")
	}

	dbSchema, envPresent := os.LookupEnv("DB_SCHEMA")
	if !envPresent {
		return nil, errors.New("DB_SCHEMA environment variable missing")
	}

	return &envVars{
		dbUserName: dbUsername,
		dbPassword: dbPassword,
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbName:     dbName,
		dbSchema:   dbSchema,
	}, nil
}

func prepareDbURL(envVars *envVars) string {
	dbURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(envVars.dbUserName, envVars.dbPassword),
		Host:   fmt.Sprintf("%s:%s", envVars.dbHost, envVars.dbPort),
		Path:   envVars.dbName,
	}

	q := dbURL.Query()
	q.Add("search_path", envVars.dbSchema)
	dbURL.RawQuery = q.Encode()
	return dbURL.String()
}
