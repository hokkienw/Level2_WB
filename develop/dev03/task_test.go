package main

import (
	"io"
	"os"
	"testing"
)

func TestSortByColumn(t *testing.T) {
	inputContent := "apple 2\nbanana 1\ncherry 3\n"
	expectedOutput := "banana 1\napple 2\ncherry 3\n"
	err := runSortTest(t, inputContent, expectedOutput, 2, false, false, false)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
}

func TestReverseSort(t *testing.T) {
	inputContent := "c\nb\na\n"
	expectedOutput := "c\nb\na\n"
	err := runSortTest(t, inputContent, expectedOutput, 1, false, true, false)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
}

func TestUniqueSort(t *testing.T) {
	inputContent := "apple\nbanana\napple\nbanana\ncherry\n"
	expectedOutput := "apple\nbanana\ncherry\n"
	err := runSortTest(t, inputContent, expectedOutput, 1, false, false, true)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
}


func runSortTest(t *testing.T, inputContent, expectedOutput string, column int, numeric, reverse, unique bool) error {
	inputFile, err := os.CreateTemp("", "input")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(inputFile.Name())

	if _, err = io.WriteString(inputFile, inputContent); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	inputFile.Close()

	outputFile, err := os.CreateTemp("", "output")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(outputFile.Name())

	err = compareSort(inputFile.Name(), outputFile.Name(), column, numeric, reverse, unique)
	if err != nil {
		return err
	}

	output, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %s", err)
	}

	if string(output) != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, string(output))
	}
	return nil
}
