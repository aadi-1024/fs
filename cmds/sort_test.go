package cmds_test

import (
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		{
			input:  "hello.go",
			output: ".go",
		},
		{
			input:  "hello.GO",
			output: ".GO",
		},
		{
			input:  "file",
			output: "",
		},
	}

	for _, v := range tt {
		actual := filepath.Ext(v.input)
		if actual != v.output {
			t.Errorf("expected: %v, got: %v\n", v.output, actual)
		}
	}
}
