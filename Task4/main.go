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
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	scanner.Scan()
	t := 0
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		var n, m int
		fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

		grid := make([][]byte, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			grid[j] = []byte(scanner.Text())
		}

		solve(grid, n, m, writer)
	}
}

func solve(grid [][]byte, n, m int, w *bufio.Writer) {
	var ax, ay, bx, by int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'A' {
				ax, ay = i, j
			} else if grid[i][j] == 'B' {
				bx, by = i, j
			}
		}
	}

	// Определяем, какой робот ближе к верхнему левому углу
	if abs(ax-0)+abs(ay-0) <= abs(bx-0)+abs(by-0) {
		// Робот A движется вверх и влево, робот B - вниз и вправо
		for ax > 0 || ay > 0 {
			if ax > 0 && grid[ax-1][ay] != '#' {
				ax--
			} else if ay > 0 && grid[ax][ay-1] != '#' {
				ay--
			}
			grid[ax][ay] = 'a'
		}

		for bx < n-1 || by < m-1 {
			if bx < n-1 && grid[bx+1][by] != '#' {
				bx++
			} else if by < m-1 && grid[bx][by+1] != '#' {
				by++
			}
			grid[bx][by] = 'b'
		}
	} else {
		// Робот B движется вверх и влево, робот A - вниз и вправо
		for bx > 0 || by > 0 {
			if bx > 0 && grid[bx-1][by] != '#' {
				bx--
			} else if by > 0 && grid[bx][by-1] != '#' {
				by--
			}
			grid[bx][by] = 'b'
		}

		for ax < n-1 || ay < m-1 {
			if ax < n-1 && grid[ax+1][ay] != '#' {
				ax++
			} else if ay < m-1 && grid[ax][ay+1] != '#' {
				ay++
			}
			grid[ax][ay] = 'a'
		}
	}

	for i := 0; i < n; i++ {
		w.WriteString(string(grid[i]) + "\n")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
