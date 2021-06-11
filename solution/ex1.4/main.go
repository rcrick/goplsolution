package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run main.go 1.txt ss
func main() {
	counts := make(map[string]int)
	lineToFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			defer f.Close()
			countLines(f, counts, lineToFiles)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, lineToFiles[line])
		}
	}
}

func addFileInfo(fileName, line string, lineToFiles map[string][]string) {
	for _, name := range lineToFiles[line] {
		if name == fileName {
			return
		}
	}
	lineToFiles[line] = append(lineToFiles[line], fileName)
}

func countLines(f *os.File, counts map[string]int, lineToFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		addFileInfo(f.Name(), input.Text(), lineToFiles)
	}
}
