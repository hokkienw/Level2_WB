package main

import (
	"bytes"
	"testing"
	"bufio"
)

func TestGrepScanner(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		pattern     string
		ignoreCase  bool
		invert      bool
		fixed       bool
		lineNum     bool
		count       bool
		before      int
		after       int
		context     int
		expectedOut string
	}{
		{
			name:        "Simple Match",
			input:       "line1\nline2\nline3\n",
			pattern:     "line",
			expectedOut: "line1\nline2\nline3\n",
		},
		{
			name:        "Match with Line Numbers",
			input:       "line1\nline2\nline3\n",
			pattern:     "line",
			lineNum:     true,
			expectedOut: "1:line1\n2:line2\n3:line3\n",
		},
		{
			name:        "Match with Context",
			input:       "line1\nline2\nline3\n",
			pattern:     "line",
			context:     1,
			expectedOut: "line1\nline2\nline3\n",
		},
		{
			name:        "Invert Match",
			input:       "line1\nline2\nline3\n",
			pattern:     "line",
			invert:      true,
			expectedOut: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdin := bytes.NewBufferString(tc.input)
			output := &bytes.Buffer{}
			scanner := bufio.NewScanner(stdin)
			grepScanner(scanner, tc.pattern, tc.ignoreCase, tc.invert, tc.fixed, tc.lineNum, tc.count, tc.before, tc.after, tc.context)

			if output.String() != tc.expectedOut {
				t.Errorf("Expected output %q, but got %q", tc.expectedOut, output.String())
			}
		})
	}
}

func TestMatch(t *testing.T) {
	testCases := []struct {
		name       string
		line       string
		pattern    string
		ignoreCase bool
		invert     bool
		fixed      bool
		expected   bool
	}{
		{
			name:     "Simple Match",
			line:     "line1",
			pattern:  "line",
			expected: true,
		},
		{
			name:     "Ignore Case Match",
			line:     "Line1",
			pattern:  "line",
			expected: true,
		},
		{
			name:     "Fixed Match",
			line:     "line1",
			pattern:  "line",
			fixed:    true,
			expected: true,
		},
		{
			name:     "Invert Match",
			line:     "line1",
			pattern:  "line",
			invert:   true,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := match(tc.line, tc.pattern, tc.ignoreCase, tc.invert, tc.fixed)
			if actual != tc.expected {
				t.Errorf("Expected match %v for line %s with pattern %s, but got %v", tc.expected, tc.line, tc.pattern, actual)
			}
		})
	}
}
