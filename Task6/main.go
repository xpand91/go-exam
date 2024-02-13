package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("tests/2")
	run(inputFile, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(r)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())

		matrix := make([][]int, n)
		for j := 0; j < n; j++ {
			matrix[j] = make([]int, m)
			scanner.Scan()
			for k, c := range scanner.Text() {
				matrix[j][k] = int(c - '0')
			}
		}

		row, col := solve(w, matrix)
		fmt.Fprintf(w, "%d %d\n", row+1, col+1)
	}
}

func solve(w io.Writer, matrix [][]int) (int, int) {
	n := len(matrix)
	m := len(matrix[0])
	minVal := 6
	minRow, minCol := -1, -1
	rowCount := make([]int, n)
	colCount := make([]int, m)

	// Найти минимальное значение и подсчитать количество в каждой строке и столбце
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] < minVal {
				minVal = matrix[i][j]
				minRow, minCol = i, j
				rowCount = make([]int, n)
				colCount = make([]int, m)
				rowCount[i]++
				colCount[j]++
			} else if matrix[i][j] == minVal {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	//fmt.Fprintf(w, "n, m: %d, %d\nminVal: %d\nminRow: %d\nminCol: %d\nrowCount(minRow): %d\ncolCount(minCol): %d\n", n, m, minVal, minRow, minCol, rowCount[minRow], colCount[minCol])

	// Если есть несколько строк или столбцов с минимальным значением, найти строку и столбец с наибольшим количеством минимальных значений и удалить их.
	maxCount := -1
	minRow, minCol = -1, -1
	for i, count := range rowCount {
		if count > maxCount {
			maxCount = count
			minRow = i
		}
	}
	for j, count := range colCount {
		if count > maxCount {
			maxCount = count
			minCol = j
			minRow = -1
		}
	}

	minValSecond := 6
	minRowSecond, minColSecond := -1, -1

	// Повторно найти минимальное значение и подсчитать количество в каждой строке и столбце
	for i := 0; i < n; i++ {
		if minRow != -1 && minRow == i {
			continue
		}
		for j := 0; j < m; j++ {
			if minCol != -1 && minCol == j {
				continue
			}
			if matrix[i][j] < minValSecond {
				minValSecond = matrix[i][j]
				minRowSecond, minColSecond = i, j
				rowCount = make([]int, n)
				colCount = make([]int, m)
				rowCount[i]++
				colCount[j]++
			} else if matrix[i][j] == minValSecond {
				rowCount[i]++
				colCount[j]++
			}
		}
	}
	//fmt.Fprintf(w, "n, m: %d, %d\nminRow: %d\nminCol: %d\nminValSecond: %d\nminRowSecond: %d\nminColSecond: %d\nrowCount(minRowSecond): %d\ncolCount(minColSecond): %d\n", n, m, minRow, minCol, minValSecond, minRowSecond, minColSecond, rowCount[minRowSecond], colCount[minColSecond])

	if minRow == -1 {
		return minRowSecond, minCol
	} else {
		return minRow, minColSecond
	}
}
