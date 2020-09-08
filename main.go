package main

import (
	"github.com/faruoqi/evento-events/db"
	"github.com/faruoqi/evento-events/rest"
)

const (
	ENDPOINT     = ":8080"
	DBCONNECTION = "mongodb://192.168.159.128:27017"
)

func main() {

	dbHandler, err := db.NewMongoDbLayer(DBCONNECTION)
	if err != nil {
		panic(err.Error())
	}
	rest.ServeAPI(ENDPOINT, dbHandler)

}
