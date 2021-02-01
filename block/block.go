package block

//block has structure block->result->[]transaction->value

type Transaction struct {
	Value string `json:"value"`
}

type Result struct {
	Transactions []Transaction `json:"transactions"`
}

type Block struct {
	Result Result `json:"result"`
}

type Stats struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}
