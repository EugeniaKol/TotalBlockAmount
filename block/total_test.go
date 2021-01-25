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
					Transactions: []Transaction{{
						Value: "0x13",
					}, {
						Value: "0x54",
					}},
				},
			},
			expected: Stats{
				Transactions: 2,
				Amount:       1.03e-16,
			},
			exError: nil,
		},

		{
			name: "Basic#2",
			args: Block{
				Result: Result{
					Transactions: []Transaction{{
						Value: "0x00",
					}, {
						Value: "0x57cf",
					}, {
						Value: "0x847",
					}},
				},
			},
			expected: Stats{
				Transactions: 3,
				Amount:       2.4597999999999998e-14,
			},
			exError: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.args.CalculateTotal()
			if res != testCase.expected {
				t.Error("expected ", testCase.expected, " got ", res)
			} else if testCase.exError != err {
				t.Error("unexpected error:", err)
			}
		})
	}
}
