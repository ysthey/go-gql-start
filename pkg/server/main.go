package server

import (
	"log"
	"path"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ysthey/go-gql-start/internal/handlers"
	"github.com/ysthey/go-gql-start/internal/orm"
	"github.com/ysthey/go-gql-start/pkg/config"
)

var host, port, gqlPath, gqlPgPath,staticDir string
var isPgEnabled, isDebug bool

func init() {
	host = config.MustGet("server-host")
	port = config.MustGet("server-port")
	staticDir = config.MustGet("server-static-dir")
	gqlPath = config.MustGet("gql-path")
	gqlPgPath = config.MustGet("gql-playground-path")
	isPgEnabled = config.MustGetBool("gql-playground-enabled")
	isDebug = config.MustGet("gin-mode") == "debug"

}

// Run web server
func Run(orm *orm.ORM) {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	if isDebug {
		r.Use(cors.Default())
	}
	r.Use(static.Serve("/", static.LocalFile(staticDir, true)))
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File(path.Join(staticDir, "index.html"))
		} else {
			c.File(path.Join(staticDir, dir, file))
		}
	})

	// Handlers
	g := r.Group("/api")
	{
		// Simple keep-alive/ping handler
		g.GET("/ping", handlers.Ping())
		g.POST("/ping", handlers.Ping())
	}

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	// pass orm to graphqlHandler
	h := handlers.GraphqlHandler(orm)
	r.POST(gqlPath, h)
	r.GET(gqlPath, h)
	log.Println("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(host + ":" + port))
}
