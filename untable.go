package textfmt

import (
	"fmt"
	"io"
	"regexp"
)

// UntableOptions contains the options for undoing a table.
type UntableOptions struct {
	MinSep int // the minimum number of spaces used to separate columns
}

// Untable reads a table from the given Reader and returns the items.
func Untable(r io.Reader, opts UntableOptions) ([]string, error) {
	if opts.MinSep < 0 {
		panic("minimum column separation must be non-negative")
	}

	split := regexp.MustCompile(fmt.Sprintf(" {%d,}", opts.MinSep))

	rows, err := readLines(r)
	if err != nil {
		return nil, fmt.Errorf("read: %v", err)
	}

	nCols := 0
	nItems := 0
	splitRows := make([][]string, 0, len(rows))
	for _, row := range rows {
		splitRow := split.Split(row, -1)
		if cols := len(splitRow); cols > nCols {
			nCols = cols
		}
		nItems += len(splitRow)
		splitRows = append(splitRows, splitRow)
	}

	items := make([]string, 0, nItems)
	for j := 0; j < nCols; j++ {
		for _, row := range splitRows {
			if j < len(row) {
				items = append(items, row[j])
			}
		}
	}
	return items, nil
}
