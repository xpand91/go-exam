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

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}

		f := make([]int, n+1)
		g := make([]int, n+1)
		h := make([]int, n+1)
		res := make([]int, n+1)

		for i := 1; i <= n; i++ {
			f[i] = 1
			if i > 1 && a[i-1] < a[i] {
				f[i] = max(f[i], g[i-1]+1)
			}
			if i > 1 && a[i-1] > a[i] {
				g[i] = max(g[i], f[i-1]+1)
				g[i] = max(g[i], h[i-1]+1)
			}
			h[i] = max(h[i-1], f[i-1]+1)
			h[i] = max(h[i], g[i-1]+1)
			res[f[i]] = max(res[f[i]], (f[i]+1)/2)
			res[g[i]] = max(res[g[i]], (g[i]+1)/2)
		}

		for i := n - 1; i >= 1; i-- {
			res[i] = max(res[i], res[i+1])
		}

		for i := 1; i <= n; i++ {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
