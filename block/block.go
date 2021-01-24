package block

type Block struct {
	Result struct {
		Transactions []struct {
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}

type Stats struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}
