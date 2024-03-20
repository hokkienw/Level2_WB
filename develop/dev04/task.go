package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	wordSet := make(map[string]bool)
	for _, word := range words {
		word = strings.ToLower(word)
		wordSet[word] = true
	}


	for word := range wordSet {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	for key, value := range anagramMap {
		if len(value) == 1 {
			delete(anagramMap, key)
		} else {
			sort.Strings(anagramMap[key])
		}
	}
	return anagramMap
}

func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)
	for key, value := range anagrams {
		fmt.Println(key+":", value)
	}
}
