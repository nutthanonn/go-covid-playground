package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nutthanonn/covid-data/cmd/infrastructure/datastore"
	"github.com/nutthanonn/covid-data/cmd/infrastructure/routers"
)

func main() {
	app := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := datastore.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	quickStart := client.Database("covid-data")
	api := app.Group("/api")
	{
		routers.CallCovid(api, quickStart)
	}

	app.Run(":8080")
}
