package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
)

type Block struct{
	Result struct{
		Transactions []struct{
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}

type Stats struct{
	Transactions int `json:"transactions"`
	Amount float64 `json:"amount"`
}

func productsHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNum, err := strconv.Atoi(vars["blockNumber"])
	request := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%x&boolean=true&apikey=YourApiKeyToken", blockNum)
	//fmt.Fprint(rw, request)
	resp, err := http.Get(request)
	//fmt.Fprint(rw, fmt.Sprintf("%X", blockNum))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	//fmt.Println(len(body))
	//fmt.Fprint(rw, string(body))
	if err != nil {
		log.Fatalln(err)
	}

	var block Block

	//err = json.Unmarshal(body, &block)

	err = json.NewDecoder(resp.Body).Decode(&block)

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(block)
	var res Stats

	for _, item := range block.Result.Transactions{
		res.Transactions++
		val := item.Value[2:]


		eth, _ := new(big.Int).SetString(val, 16)
		f := new(big.Float).SetInt(eth)
		wei, acc := f.Float64()
		fmt.Println(acc)

		//eth, err := strconv.ParseUint(val, 16, 64)
		//var wei, ok = new(big.Int).SetString(val, 16)
		if  err!=nil {
			log.Fatalln(err)
		}
		//res.Amount += float64(eth)/math.Pow(10, 18 - float64(digits))
		res.Amount += wei/math.Pow(10, 18)
		//fmt.Fprint(rw, res, "\n")
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	_ = json.NewEncoder(rw).Encode(res)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/block/{blockNumber:[0-9]+}/total", productsHandler)
	http.Handle("/",router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
