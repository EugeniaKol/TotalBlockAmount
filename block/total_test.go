package block

import "testing"

func TestBlock_CalculateTotal(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		args     Block
		expected Stats
		exError  error
	}{
		{
			name: "Basic",
			args: Block{
				Result: Result{
					Transactions: []Transaction{
						{
							Value: "0x124c6",
						},
						{
							Value: "0x65778f",
						},
					},
				},
			},
			expected: Stats{
				Transactions: 2,
				Amount:       0,
			},
			exError: nil,
		},

		{
			name: "Basic",
			args: Block{
				Result: Result{
					Transactions: []Transaction{
						{
							Value: "0x1ad3467",
						},
						{
							Value: "0x657fg",
						},
					},
				},
			},
			expected: Stats{
				Transactions: 2,
				Amount:       0,
			},
			exError: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.args.CalculateTotal()
			if res != testCase.expected {
				t.Errorf("expected total %f, got %f", testCase.expected.Amount, res.Amount)
			} else if testCase.exError != err {
				t.Error("unexpected error:", err)
			}
		})
	}
}