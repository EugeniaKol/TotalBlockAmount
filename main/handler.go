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
	blockNum, _ := strconv.Atoi(vars["blockNumber"])

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

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)
}
