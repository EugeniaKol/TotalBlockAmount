package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

type DbClient struct {
	Client *redis.Client
}

func GetRedisClient() {
	StClient.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := StClient.Client.Ping(StClient.Client.Context()).Result()
	if err != nil {
		fmt.Println(pong, err)
		Conf.EnableCaching = false
	}
}

var StClient DbClient

func main() {
	SetConfig(ConfFile)
	GetRedisClient()

	router := mux.NewRouter()

	router.HandleFunc("/api/block/{blockNumber:[0-9]+}/total", TotalHandler)
	http.Handle("/", router)
	fmt.Println("Server is listening...")

	if err := http.ListenAndServe(Conf.Port, nil); err != nil {
		fmt.Println(err)
	}
}
