package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ysthey/go-gql-start/internal/handlers"
	"github.com/ysthey/go-gql-start/pkg/config"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = config.MustGet("server-host")
	port = config.MustGet("server-port")
	gqlPath = config.MustGet("gql-path")
	gqlPgPath = config.MustGet("gql-playground-path")
	isPgEnabled = config.MustGetBool("gql-playground-enabled")
}

// Run web server
func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(host + ":" + port))
}
