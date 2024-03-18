package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	column := flag.Int("k", 0, "указание колонки для сортировки")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Ошибка при чтении файла:", err)
			os.Exit(1)
		}
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
		if err == io.EOF {
			break
		}
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return compare(lines[i], lines[j], *column, *numeric) < 0
	})

	if *reverse {
		reverseSlice(lines)
	}

	if *unique {
		lines = removeDuplicates(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func compare(a, b string, column int, numeric bool) int {
	if column > 0 {
		wordsA := strings.Fields(a)
		wordsB := strings.Fields(b)
		if column <= len(wordsA) && column <= len(wordsB) {
			a = wordsA[column-1]
			b = wordsB[column-1]
		}
	}

	if numeric {
		numA, errA := strconv.Atoi(a)
		numB, errB := strconv.Atoi(b)
		if errA == nil && errB == nil {
			if numA < numB {
				return -1
			} else if numA > numB {
				return 1
			} else {
				return 0
			}
		}
	}

	return strings.Compare(a, b)
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))
	for _, str := range slice {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}
