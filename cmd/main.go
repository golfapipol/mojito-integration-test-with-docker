package main

import (
	"mojito/todo/api"
	"mojito/todo/repository"
	"log"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	api := api.API{
		Repository: &repository.RepositoryMongo{
			Session: session,
		},
	}
	route := gin.Default()
	route.GET("api/v1/country", api.ViewCountry)
	route.Run(":3300")
}
