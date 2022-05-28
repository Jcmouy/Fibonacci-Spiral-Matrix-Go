package main

import (
	_ "fibonacci-spiral-matrix-go/docs"
	"fibonacci-spiral-matrix-go/internal/api"
	"fibonacci-spiral-matrix-go/internal/config/database"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"log"
	"os"
)

// @title Fibonacci Spiral Matrix API
// @version 1.0.0
// @description This is the fibonacci spiral matrix restful api server.
// @host localhost:8080
// @BasePath /
func main() {
	app, err := api.NewAppServer()

	if err != nil {
		log.Fatalf("AppServer Error: %s", err.Error())
	}
	defer app.Close()

	err = database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&user.User{})

	if err := app.Start(); err != nil {
		log.Printf("%s", err.Error())
		os.Exit(1)
	}
}
