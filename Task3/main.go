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
		var s string
		fmt.Fscan(in, &s)
		if isValid(s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func isValid(s string) bool {
	m, c, d := 0, 0, 0
	prev := ' '

	if s[0] != 'M' {
		return false
	}

	for _, curr := range s {
		switch curr {
		case 'M':
			if prev != 'C' && prev != 'D' && prev != ' ' {
				return false
			}
			m++
		case 'R':
			if prev != 'M' {
				return false
			}
		case 'C':
			if prev != 'M' && prev != 'R' {
				return false
			}
			c++
		case 'D':
			if prev != 'M' || d != m-c-1 {
				return false
			}
			d++
		default:
			return false
		}
		prev = curr
	}

	return prev == 'D'
}
