package block

type Transaction struct {
	Value string `json:"value"`
}

type Result struct {
	Transactions []Transaction `json:"transactions"`
}

type Block struct {
	Result Result `json:"result"`
}

/*type Block struct {
	Result struct {
		Transactions []struct {
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}*/

type Stats struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}
