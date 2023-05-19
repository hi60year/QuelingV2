package main

import (
	"encoding/json"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/db"
	"main/graph"
	"os"
)

const defaultPort = "8080"

func main() {
	confFile, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := confFile.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := json.NewDecoder(confFile)
	configuration := make(map[string]interface{})

	if err = decoder.Decode(&configuration); err != nil {
		panic(err)
	}

	router := gin.Default()

	var client *mongo.Client

	if client, err = db.Connect(configuration["db-connection-string"].(string)); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	defer db.Disconnect(client)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(client)}))
	// srv.Use(extension.FixedComplexityLimit(20))

	router.POST("/query", gin.WrapH(srv))

	// Start the server on port 8080
	if err = router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
