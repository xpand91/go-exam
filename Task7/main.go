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

	var n, m int
	fmt.Fscan(in, &n)

	logins := make(map[string]bool)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		logins[s] = true
		for j := 0; j < len(s)-1; j++ {
			logins[s[:j]+string(s[j+1])+string(s[j])+s[j+2:]] = true
		}
	}

	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		var t string
		fmt.Fscan(in, &t)
		if logins[t] {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}
