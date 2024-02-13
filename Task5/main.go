package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	t := scanner.Text()
	testCases := atoi(t)

	for i := 0; i < testCases; i++ {
		scanner.Scan()
		n := atoi(scanner.Text())
		var jsonLines []string
		for j := 0; j < n; j++ {
			scanner.Scan()
			jsonLines = append(jsonLines, scanner.Text())
		}
		jsonData := strings.Join(jsonLines, "")
		var root Folder
		err := json.Unmarshal([]byte(jsonData), &root)
		if err != nil {
			fmt.Fprintln(w, "Error parsing JSON:", err)
			return
		}
		infectedFiles := countInfectedFiles(root, false)
		fmt.Fprintln(w, infectedFiles)
	}
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func countInfectedFiles(folder Folder, isInfected bool) int {
	count := 0
	for _, file := range folder.Files {
		if strings.HasSuffix(file, ".hack") {
			isInfected = true
			break
		}
	}
	if isInfected {
		count += len(folder.Files)
	}
	for _, subfolder := range folder.Folders {
		count += countInfectedFiles(subfolder, isInfected)
	}
	return count
}
