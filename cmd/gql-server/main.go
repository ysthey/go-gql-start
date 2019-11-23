package main

import (
	"log"

	"github.com/ysthey/go-gql-start/internal/orm"
	"github.com/ysthey/go-gql-start/pkg/server"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}
