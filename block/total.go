package block

import (
	"errors"
	"math"
	"math/big"
)

func (block Block) CalculateTotal() (Stats, error) {
	var res Stats
	for _, item := range block.Result.Transactions {
		res.Transactions++
		val := item.Value[2:]

		num, _ := new(big.Int).SetString(val, 16)
		if num == nil {
			return Stats{}, errors.New("cant parse value")
		}

		f := new(big.Float).SetInt(num)
		wei, _ := f.Float64()
		res.Amount += wei / math.Pow(10, 18)
	}

	return res, nil
}
