package pprint

import (
	"bufio"
	"fmt"
	"strings"
)

func PrintTable[T string | int](t [][]T, w *bufio.Writer) {
	for _, row := range t {
		vals := ""
		for _, v := range row {
			vals += fmt.Sprintf("%v ", v)
		}
		_, _ = w.WriteString(strings.TrimSpace(vals) + "\n")
	}
}

func PrintSlice[T string | int](s []T, w *bufio.Writer) {
	vals := ""
	for _, v := range s {
		vals += fmt.Sprintf("%v ", v)
	}
	_, _ = w.WriteString(strings.TrimSpace(vals) + "\n")
}
