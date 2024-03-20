package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestSortLines(t *testing.T) {
	testCases := []struct {
		input    string
		args     []string
		expected string
	}{
		{
			input: `banana
apple
orange`,
			args:     []string{"-k", "0"},
			expected: "apple\nbanana\norange\n",
		},
		{
			input: `5 apple
10 banana
3 orange`,
			args:     []string{"-k", "1", "-n"},
			expected: "3 orange\n5 apple\n10 banana\n",
		},
		{
			input: `orange
banana
apple`,
			args:     []string{"-k", "0", "-r"},
			expected: "orange\nbanana\napple\n",
		},

		{
			input: `apple
banana
apple`,
			args:     []string{"-k", "0", "-u"},
			expected: "apple\nbanana\n",
		},
	}

	for _, tc := range testCases {
		tmpfile, err := ioutil.TempFile("", "test_input")
		if err != nil {
			t.Fatal("Error creating temporary file:", err)
		}
		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write([]byte(tc.input)); err != nil {
			t.Fatal("Error writing to temporary file:", err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatal("Error closing temporary file:", err)
		}
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()

		os.Stdout = oldStdout

		var out bytes.Buffer
		out.ReadFrom(r)

		actual := out.String()
		if actual != tc.expected {
			t.Errorf("Unexpected output for args %v: expected %q, got %q", tc.args, tc.expected, actual)
		}
	}
}
