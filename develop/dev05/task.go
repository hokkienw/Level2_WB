package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines before match")
	context := flag.Int("C", 0, "Print ±N lines around match")
	count := flag.Bool("c", false, "Print only the count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: grep [options] pattern [file ...]")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	if len(files) == 0 {
		grepStdin(pattern, *ignoreCase, *invert, *fixed, *lineNum, *count)
	} else {
		for _, file := range files {
			grepFile(file, pattern, *ignoreCase, *invert, *fixed, *lineNum, *count, *before, *after, *context)
		}
	}
}

func grepStdin(pattern string, ignoreCase, invert, fixed, lineNum, count bool) {
	scanner := bufio.NewScanner(os.Stdin)
	grepScanner(scanner, pattern, ignoreCase, invert, fixed, lineNum, count, 0, 0, 0)
}

func grepFile(filename, pattern string, ignoreCase, invert, fixed, lineNum, count bool, before, after, context int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grepScanner(scanner, pattern, ignoreCase, invert, fixed, lineNum, count, before, after, context)
}

func grepScanner(scanner *bufio.Scanner, pattern string, ignoreCase, invert, fixed, lineNum, count bool, before, after, context int) {
	var buffer []string
	var matchIndex int

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		if match(line, pattern, ignoreCase, invert, fixed) {
			matchIndex++
			if count {
				continue
			}
			printLine(line, lineNum, matchIndex)
			if after > 0 {
				for j := 0; j < after && scanner.Scan(); j++ {
					i++
					line = scanner.Text()
					printLine(line, lineNum, i)
				}
			}
			buffer = nil
		} else {
			if before > 0 || context > 0 {
				buffer = append(buffer, line)
				if len(buffer) > before+context {
					buffer = buffer[1:]
				}
			}
		}
	}
}

func match(line, pattern string, ignoreCase, invert, fixed bool) bool {
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if fixed {
		return strings.Contains(line, pattern)
	}
	if invert {
		return !strings.Contains(line, pattern)
	}
	return strings.Contains(line, pattern)
}

func printLine(line string, lineNum bool, lineNumber int) {
	if lineNum {
		fmt.Printf("%d:", lineNumber)
	}
	fmt.Println(line)
}
