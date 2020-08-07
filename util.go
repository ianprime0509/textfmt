package textfmt

import (
	"bufio"
	"io"
	"strings"
)

func readLines(r io.Reader) ([]string, error) {
	var items []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		item := strings.TrimSpace(s.Text())
		if item != "" {
			items = append(items, item)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
