package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type config struct {
	serverPortNumber string
	serverAddress    string
	redisPortNumber  string
	redisHostname    string
	redisPassword    string
	redisDB          int
}

var (
	defaultConfig = &config{
		serverPortNumber: "3000",
		serverAddress:    "0.0.0.0",
		redisPortNumber:  "6379",
		redisHostname:    "localhost",
		redisPassword:    "",
		redisDB:          0,
	}
)

type application struct {
	ctx    context.Context
	rdb    *redis.Client
	config *config
}

func main() {

	app := &application{
		config: defaultConfig,
	}

	app.setup()

	app.dbConnect()

	app.start()

}
