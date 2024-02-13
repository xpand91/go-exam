package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscanf(r, "%d\n", &t)

	for i := 0; i < t; i++ {
		var n int
		var p float64
		fmt.Fscanf(r, "%d %f\n", &n, &p)

		var totalLost float64
		for j := 0; j < n; j++ {
			var ai int
			fmt.Fscanf(r, "%d\n", &ai)

			commission := float64(ai) * p / 100
			lost := commission - float64(int(commission))
			totalLost += lost
		}

		fmt.Fprintf(w, "%.2f\n", totalLost)
	}
}
