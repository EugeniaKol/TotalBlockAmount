package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	SetConfig(ConfFile)
	router := mux.NewRouter()
	router.HandleFunc("/api/block/{blockNumber:[0-9]+}/total", TotalHandler)
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(Conf.Port, nil); err != nil {
		fmt.Println(err)
	}
}
