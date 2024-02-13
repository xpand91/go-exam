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
	var a, b int
	fmt.Fscanf(r, "%d %d", &a, &b)
	fmt.Fprintln(w, a-b)
}
