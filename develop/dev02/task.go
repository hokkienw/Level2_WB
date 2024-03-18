package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var escape bool

	for _, char := range s {
		if unicode.IsDigit(char) && !escape {
			count, _ := strconv.Atoi(string(char))
			if count == 0 {
				return "", errors.New("некорректная строка")
			}
			result.WriteString(strings.Repeat(string(prev), count-1))
		} else if char == '\\' && !escape {
			escape = true
		} else {
			if escape {
				result.WriteRune(char)
				escape = false
			} else {
				result.WriteRune(char)
				prev = char
			}
		}
	}

	return result.String(), nil
}

func main() {
	stringsToUnpack := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, s := range stringsToUnpack {
		unpacked, err := UnpackString(s)
		if err != nil {
			println("Ошибка:", err.Error())
		} else {
			println("Распакованная строка:", unpacked)
		}
	}
}

