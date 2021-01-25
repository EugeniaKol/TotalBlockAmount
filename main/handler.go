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

	if Conf.EnableCaching == true {
		if isCached, value := CacheGetRequest(blockStr); isCached == true {
			rw.Header().Set("content-type", "application/json")
			rw.WriteHeader(200)
			fmt.Fprint(rw, value)
			return
		}
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

	if Conf.EnableCaching == true {
		if err = CachePostRequest(blockStr, cacheInput); err != nil {
			fmt.Println(err)
			return
		}
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)
}
