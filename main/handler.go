package main

import (
	"encoding/json"
	"fmt"
	tools "github.com/EugeniaKol/forums_system/server/tools"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func TotalHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockStr := vars["blockNumber"]
	blockNum, _ := strconv.Atoi(blockStr)

	if Conf.EnableCaching == true {
		if isCached, value := CacheGetRequest(blockStr); isCached == true {
			tools.WriteJSONOk(rw, value)
			return
		}
	}

	//fmt.Println("not cashed")
	block, err := BlockRequest(blockNum)
	if err != nil {
		fmt.Println(err)
		tools.WriteJSONBadRequest(rw, err.Error())
		return
	}

	res, err := block.CalculateTotal()
	if err != nil {
		fmt.Println(err)
		tools.WriteJSONInternalError(rw)
		return
	}

	cacheInput, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		tools.WriteJSONInternalError(rw)
	}

	if Conf.EnableCaching == true {
		if err = CachePostRequest(blockStr, cacheInput); err != nil {
			fmt.Println(err)
			return
		}
	}

	tools.WriteJSONOk(rw, res)
	/*rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)*/
}
