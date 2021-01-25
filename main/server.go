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

func GetRedisClient() DbClient {
	var client DbClient
	client.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

var StClient DbClient

func main() {
	SetConfig(ConfFile)
	StClient = GetRedisClient()
	pong, err := StClient.Client.Ping(StClient.Client.Context()).Result()
	fmt.Println(pong, err)
	router := mux.NewRouter()
	router.HandleFunc("/api/block/{blockNumber:[0-9]+}/total", TotalHandler)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(Conf.Port, nil); err != nil {
		fmt.Println(err)
	}
}
