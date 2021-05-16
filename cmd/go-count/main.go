package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	portNumber string = "3000"
)

type application struct {
	ctx context.Context
	rdb *redis.Client
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	app := &application{
		ctx: context.Background(),
		rdb: rdb,
	}

	app.setup()

	app.start()

}
