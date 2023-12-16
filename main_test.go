package main

import (
	"testing"
)

func TestIsValidChunk(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{
			input:  "[(<>})]",
			output: 4,
		},
		{
			input:  "(>",
			output: 1,
		},
		{
			input:  "[[[[[[[]]]]]]({<>})]",
			output: -1,
		},
		{
			input:  "<{[()}>]",
			output: 5,
		},
		{
			input:  "{[<(<>)]>}",
			output: 7,
		},
		{
			input:  "(<[{<>}]>)",
			output: -1,
		},
		{
			input:  "[{(<[()]>)<>}]",
			output: -1,
		},
		{
			input:  "[{([]<>)}([]<>)]",
			output: -1,
		},
		{
			input:  "[<<>]",
			output: 4,
		},
		{
			input:  "({}<>])",
			output: 5,
		},
		{
			input:  "()",
			output: -1,
		},
		{
			input:  "{}{}",
			output: -1,
		},
		{
			input:  "{}<()[]>",
			output: -1,
		},
		{
			input:  "{}<>]",
			output: 4,
		},
		{
			input:  "(<{()}{}{}><<>{[]}>[<<>><{}>{}>[]<>][]{}[])",
			output: 30,
		},
	}

	for _, c := range cases {
		t.Run("isValidChunk", func(t *testing.T) {
			got := isValidChunk(c.input)
			if got != c.output {
				t.Errorf("expected %v got %v", c.output, got)
			}
		})
	}
}
