package textfmt

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mattn/go-runewidth"
)

// TableOptions contains the options for formatting a table.
type TableOptions struct {
	Columns int // the number of columns to include in the table
	MinSep  int // the minimum number of spaces to use to separate columns
	Width   int // the ideal width of the table
}

// Table formats the given items into a tabular structure and writes the results
// to the given Writer. The provided options control the formatting of the
// table.
func Table(items []string, w io.Writer, opts TableOptions) error {
	if opts.Columns <= 0 {
		panic("number of columns must be positive")
	} else if opts.MinSep < 0 {
		panic("minimum column separation must be non-negative")
	}

	maxColWidth := 0
	totalColWidth := 0
	colWidths := make([]int, opts.Columns)
	nRows := len(items) / opts.Columns
	extraItems := len(items) % opts.Columns
	if extraItems > 0 {
		nRows++
	}

	for j := range colWidths {
		maxW := 0
		for i := 0; i < nRows; i++ {
			itemIdx := itemIndex(i, j, nRows, extraItems)
			if itemIdx < 0 || itemIdx >= len(items) {
				break
			}
			if w := runewidth.StringWidth(items[itemIdx]); w > maxW {
				maxW = w
			}
		}
		colWidths[j] = maxW
		if maxW > maxColWidth {
			maxColWidth = maxW
		}
		totalColWidth += maxW
	}

	extraSpace := opts.Width - totalColWidth - opts.MinSep*(len(colWidths)-1)
	if extraSpace > 0 {
		opts.MinSep += extraSpace / (len(colWidths) - 1)
	}
	spaceString := strings.Repeat(" ", maxColWidth+opts.MinSep)

	bw := bufio.NewWriter(w)
	for i := 0; i < nRows; i++ {
		for j, colW := range colWidths {
			itemIdx := itemIndex(i, j, nRows, extraItems)
			if itemIdx < 0 || itemIdx >= len(items) {
				break
			}

			item := items[itemIdx]
			if _, err := bw.WriteString(item); err != nil {
				return fmt.Errorf("write: %v", err)
			}

			var needSeparator bool
			if i < nRows-1 || extraItems == 0 {
				needSeparator = j < len(colWidths)-1
			} else {
				needSeparator = j < extraItems-1
			}
			if needSeparator {
				spaces := colW - runewidth.StringWidth(item) + opts.MinSep
				if _, err := bw.WriteString(spaceString[:spaces]); err != nil {
					return fmt.Errorf("write: %v", err)
				}
			}
		}
		if _, err := bw.WriteRune('\n'); err != nil {
			return fmt.Errorf("write: %v", err)
		}
	}
	if err := bw.Flush(); err != nil {
		return fmt.Errorf("write: %v", err)
	}

	return nil
}

// TableLines reads all lines from the given Reader and formats them into a table
// using Table.
func TableLines(r io.Reader, w io.Writer, opts TableOptions) error {
	items, err := readLines(r)
	if err != nil {
		return fmt.Errorf("read: %v", err)
	}
	return Table(items, w, opts)
}

func itemIndex(row, col, nRows, extraItems int) int {
	if extraItems == 0 || col < extraItems {
		return nRows*col + row
	} else if row == nRows-1 {
		return -1
	}
	return nRows*extraItems + (nRows-1)*(col-extraItems) + row
}
