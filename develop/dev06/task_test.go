package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintSelectedFields(t *testing.T) {
	line := "1\t2\t3"
	selectedFields := []string{"1", "3"}
	delimiter := "\t"
	expectedOutput := "1\t3\n"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printSelectedFields(line, selectedFields, delimiter)
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
