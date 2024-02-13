package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		var s string
		fmt.Fscan(in, &s)

		count := map[rune]int{'X': 0, 'Y': 0, 'Z': 0}
		for _, c := range s {
			count[c]++
		}

		if count['X']%2 == 0 && count['Y']%2 == 0 && count['Z']%2 == 0 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
