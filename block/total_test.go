package block

import "testing"

func TestBlock_CalculateTotal(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		args     []string
		expected Stats
		exError  error
	}{
		{
			name: "Basic",
			args: []string{"0x13", "0x54"},
			expected: Stats{
				Transactions: 2,
				Amount:       1.03e-16,
			},
			exError: nil,
		},

		{
			name: "Basic#2",
			args: []string{"0x00", "0x57cf", "0x847"},
			expected: Stats{
				Transactions: 3,
				Amount:       2.4597999999999998e-14,
			},
			exError: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			var block Block

			for _, value := range testCase.args {
				block.Result.Transactions = append(block.Result.Transactions, Transaction{Value: value})
			}

			res, err := block.CalculateTotal()

			if res != testCase.expected {
				t.Error("expected ", testCase.expected, " got ", res)
			} else if testCase.exError != err {
				t.Error("unexpected error:", err)
			}
		})
	}
}
