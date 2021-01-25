package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//b "github.com/EugeniaKol/TotalBlockAmount/block"
)

func TotalHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockStr := vars["blockNumber"]
	blockNum, _ := strconv.Atoi(blockStr)

	fmt.Println("searching block", blockStr)
	cached, err := StClient.Client.Get(StClient.Client.Context(), blockStr).Result()

	fmt.Print("result from redis for: ", blockNum, cached)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("got cashed", cached)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprint(rw, cached)
		return
	}

	//fmt.Println("not cashed")
	block, err := BlockRequest(blockNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := block.CalculateTotal()
	if err != nil {
		fmt.Println(err)
		return
	}

	cacheInput, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	err = StClient.Client.Set(StClient.Client.Context(), blockStr, cacheInput, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)
}
