package block

type Block struct {
	Result struct {
		Transactions []struct {
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}
