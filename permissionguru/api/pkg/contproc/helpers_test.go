package contproc

import "testing"

func Test_matchTexts(t *testing.T) {
	tests := []struct {
		description string
		sources     []string
		target      string
		output      string
	}{
		{
			description: "exact match returned",
			sources: []string{
				"first selection text",
				"second selection text",
			},
			target: "first selection text",
			output: "first selection text",
		},
		{
			description: "inexact match returned",
			sources: []string{
				"firstest selections text",
				"second selection text",
			},
			target: "first selection text",
			output: "firstest selections text",
		},
	}

	for _, test := range tests {
		output := matchTexts(test.sources, test.target)
		if output != test.output {
			t.Errorf("incorrect output, received: %s, expected: %s", output, test.output)
		}
	}
}
