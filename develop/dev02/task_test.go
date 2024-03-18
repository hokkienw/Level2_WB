package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{

		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
	}

	for _, testCase := range testCases {
		actual, err := UnpackString(testCase.input)

		if testCase.err && err == nil {
			t.Errorf("Ожидалась ошибка, но получено nil")
		}

		if !testCase.err && err != nil {
			t.Errorf("Ожидалось nil, но получена ошибка: %v", err)
		}

		if actual != testCase.expected {
			t.Errorf("Неправильный результат. Ожидалось '%s', получено '%s'", testCase.expected, actual)
		}
	}
}

// func main() {
// 	TestUnpackString()
// }