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
