package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"ianjohnson.xyz/textfmt"
)

var nCols = pflag.IntP("columns", "c", 2, "number of columns to use in the table")
var minSep = pflag.IntP("min-separation", "s", 2, "minimum number of spaces to separate columns")
var width = pflag.IntP("width", "w", 0, "ideal width of the table, in characters")

func main() {
	pflag.Parse()

	if err := textfmt.TableLines(os.Stdin, os.Stdout, textfmt.TableOptions{
		Columns: *nCols,
		MinSep:  *minSep,
		Width:   *width,
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
