package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if !validate(str) {
		return "", ErrInvalidString
	}

	var strBuilder strings.Builder
	for _, v := range str {
		if !unicode.IsDigit(v) {
			strBuilder.WriteString(string(v))
			continue
		}

		sbString := strBuilder.String()
		curLet := string(sbString[len(sbString)-1])
		curFactor, _ := strconv.Atoi(string(v - 1))
		if curFactor == 0 {
			str = sbString[:len(sbString)-1]
			strBuilder.Reset()
			strBuilder.WriteString(str)
			continue
		}
		curStr := strings.Repeat(curLet, curFactor)
		strBuilder.WriteString(curStr)
	}

	return strBuilder.String(), nil
}

func validate(str string) bool {
	if unicode.IsDigit(rune(str[0])) {
		return false
	}
	for i, v := range str {
		if !unicode.IsDigit(v) {
			continue
		}
		if next := string(str[i+1]); next != "" {
			if unicode.IsDigit(rune(next[0])) {
				return false
			}
		}
	}

	return true
}
