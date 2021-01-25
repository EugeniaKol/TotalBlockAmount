package main

import (
	"encoding/json"
	"errors"
	"fmt"
	b "github.com/EugeniaKol/TotalBlockAmount/block"
	"net/http"
)

func BlockRequest(blockNum int) (b.Block, error) {
	request := fmt.Sprintf(RequestString, blockNum, Conf.ApiKey)
	resp, err := http.Get(request)
	if err != nil {
		return b.Block{}, err
	}

	defer resp.Body.Close()
	var block b.Block
	err = json.NewDecoder(resp.Body).Decode(&block)

	if err != nil {
		return b.Block{}, errors.New("api key usage limit reached, please try again")
	}
	return block, nil
}

func CacheGetRequest(block string) (isCached bool, res b.Stats) {
	fmt.Println("searching block", block)
	cached, err := StClient.Client.Get(StClient.Client.Context(), block).Result()
	var stats b.Stats

	_ = json.Unmarshal([]byte(cached), &stats)
	fmt.Print("result from redis for: ", block, "  ", cached)

	if err != nil {
		fmt.Println(err)
		return false, b.Stats{}
	}
	return true, stats
}

func CachePostRequest(block string, value []byte) error {
	err := StClient.Client.Set(StClient.Client.Context(), block, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
