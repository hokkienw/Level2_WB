package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя": {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			words: []string{"apple", "banana", "orange", "kiwi"},
			expected: map[string][]string{},
		},
		{
			words: []string{"listen", "silent", "enlist"},
			expected: map[string][]string{
				"eilnst": {"enlist", "listen", "silent"},
			},
		},

		{
			words: []string{"hello"},
			expected: map[string][]string{},
		},
	}

	for _, tc := range testCases {
		actual := findAnagrams(tc.words)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Unexpected result for words %v: expected %v, got %v", tc.words, tc.expected, actual)
		}
	}
}
