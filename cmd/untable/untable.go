package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"ianjohnson.xyz/textfmt"
)

var minSep = pflag.IntP("min-separation", "s", 2, "minimum number of spaces to separate columns")

func main() {
	pflag.Parse()

	items, err := textfmt.Untable(os.Stdin, textfmt.UntableOptions{
		MinSep: *minSep,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, item := range items {
		if _, err := fmt.Println(item); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
