package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Delimiter")
	separated := flag.Bool("s", false, "Print only lines with delimiter")

	flag.Parse()

	selectedFields := strings.Split(*fields, ",")

	processInput(selectedFields, *delimiter, *separated)
}

func processInput(selectedFields []string, delimiter string, separated bool) {
	if delimiter == "" {
		delimiter = "\t"
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if separated && !strings.Contains(line, delimiter) {
			continue
		}
		printSelectedFields(line, selectedFields, delimiter)
	}
}

func printSelectedFields(line string, selectedFields []string, delimiter string) {
	fields := strings.Split(line, delimiter)

	var selected []string
	for _, fieldIndex := range selectedFields {
		index, err := strconv.Atoi(fieldIndex)
		if err != nil || index < 1 || index > len(fields) {
			continue
		}
		selected = append(selected, fields[index-1])
	}

	fmt.Println(strings.Join(selected, delimiter))
}
