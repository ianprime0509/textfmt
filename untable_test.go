package textfmt

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUntable(t *testing.T) {
	tests := []struct {
		input  string
		opts   UntableOptions
		output []string
	}{
		{
			"one  two\nthree  four\n",
			UntableOptions{MinSep: 2},
			[]string{"one", "three", "two", "four"},
		},
		{
			"one    three\n    \n\ntwo    four",
			UntableOptions{MinSep: 1},
			[]string{"one", "two", "three", "four"},
		},
	}

	for _, test := range tests {
		output, err := Untable(strings.NewReader(test.input), test.opts)
		if err != nil {
			t.Errorf("Unexpected error for input %q with options %#v: %v", test.input, test.opts, err)
			continue
		}
		if !cmp.Equal(output, test.output) {
			t.Errorf("For input %q with options %#v: want %q, got %q", test.input, test.opts, test.output, output)
		}
	}
}
