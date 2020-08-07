package textfmt

import (
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	tests := []struct {
		input  []string
		opts   TableOptions
		output string
	}{
		{
			[]string{"abc", "def", "ghi", "jkl"},
			TableOptions{Columns: 2, MinSep: 2},
			"abc  ghi\ndef  jkl\n",
		},
		{
			[]string{"one", "two", "three", "four", "five", "six", "seven"},
			TableOptions{Columns: 3, MinSep: 2},
			"one    four  six\ntwo    five  seven\nthree\n",
		},
		{
			[]string{"1", "2", "3", "four", "5", "six", "seven"},
			TableOptions{Columns: 5, MinSep: 3},
			"1   3      5   six   seven\n2   four\n",
		},
	}

	for _, test := range tests {
		w := new(strings.Builder)
		if err := Table(test.input, w, test.opts); err != nil {
			t.Errorf("Unexpected error for input %q with options %#v: %v", test.input, test.opts, err)
			continue
		}
		if output := w.String(); output != test.output {
			t.Errorf("For input %q with options %#v: want %q, got %q", test.input, test.opts, test.output, output)
		}
	}
}
