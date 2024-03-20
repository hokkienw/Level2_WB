package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fileString struct {
	Line string
	Key  string
}

func main() {
	columnNumber := flag.Int("c", 1, "column number for sorting")
	isNumeric := flag.Bool("n", false, "sort by numeric value")
	isReverse := flag.Bool("r", false, "reverse the sorting order")
	isUnique := flag.Bool("u", false, "output only the first of an equal run")
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalln("incorrect input")
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	err := compareSort(inputFile, outputFile, *columnNumber, *isNumeric, *isReverse, *isUnique)
	if err != nil {
		log.Fatalln(err)
	}
}

func compareSort(inputFile, outputFile string, columnNumber int, isNumeric, isReverse, isUnique bool) error {
	lines, err := fileReader(inputFile)
	if err != nil {
		return err
	}

	var data []fileString
	for _, line := range lines {
		columns := strings.Fields(line)
		var key string
		if columnNumber-1 < len(columns) {
			key = columns[columnNumber-1]
		}
		data = append(data, fileString{line, key})
	}

	less := func(i, j int) bool {
		if isNumeric {
			numI, errI := strconv.ParseFloat(data[i].Key, 64)
			numJ, errJ := strconv.ParseFloat(data[j].Key, 64)
			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}
		return data[i].Key < data[j].Key
	}

	if isReverse {
		sort.SliceStable(data, func(i, j int) bool {
			return less(j, i)
		})
	} else {
		sort.SliceStable(data, less)
	}

	if isUnique {
		data = removeDuplicates(data)
	}

	if err := fileWriter(data, outputFile); err != nil {
		return err
	}
	return nil
}

func fileReader(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func fileWriter(data []fileString, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range data {
		fmt.Fprintln(w, line.Line)
	}
	return w.Flush()
}

func removeDuplicates(data []fileString) []fileString {
	if len(data) == 0 {
		return data
	}

	uniqueData := []fileString{data[0]}
	lastLine := data[0].Line
	for _, d := range data[1:] {
		if d.Line != lastLine {
			uniqueData = append(uniqueData, d)
			lastLine = d.Line
		}
	}
	return uniqueData
}
