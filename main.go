package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(r *bufio.Reader, w *bufio.Writer) {

}

func main() {
	var t int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		solve(r, w)
	}
	_ = w.Flush()
}
