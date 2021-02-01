package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ErrorResp struct {
	Message string `json:"message"`
}

func TotalHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	blockStr := vars["blockNumber"]
	if blockStr == "" {
		err := "no block number specified"
		fmt.Println(err)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(400)
		_ = json.NewEncoder(rw).Encode(ErrorResp{"bad request: " + err})
		return
	}

	blockNum, err := strconv.Atoi(blockStr)
	if err != nil {
		err := "incorrect block number input"
		fmt.Println(err)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(400)
		_ = json.NewEncoder(rw).Encode(ErrorResp{"bad request: " + err})
		return
	}

	//if caching is enabled
	if Conf.EnableCaching == true {
		if isCached, value := CacheGetRequest(blockStr); isCached == true {
			rw.Header().Set("content-type", "application/json")
			rw.WriteHeader(200)
			_ = json.NewEncoder(rw).Encode(value)
			return
		}
	}

	//if caching is disabled
	block, err := BlockRequest(blockNum)
	if err != nil {
		fmt.Println(err)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(400)
		_ = json.NewEncoder(rw).Encode(ErrorResp{err.Error()})
		return
	}

	res, err := block.CalculateTotal()
	if err != nil {
		fmt.Println(err)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(500)
		_ = json.NewEncoder(rw).Encode(ErrorResp{err.Error()})
		return
	}

	cacheInput, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(500)
		_ = json.NewEncoder(rw).Encode(ErrorResp{err.Error()})
	}

	if Conf.EnableCaching == true {
		if err = CachePostRequest(blockStr, cacheInput); err != nil {
			fmt.Println(err)
		}
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)
}
