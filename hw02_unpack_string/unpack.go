package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if regexpValidate(str) != nil {
		return "", regexpValidate(str)
	}

	splitByRegexp := splitByRegex(str)

	var strBuilder strings.Builder
	var curStr string
	for _, v := range splitByRegexp {
		if string(v[0]) == "\\" {
			if len(v) == 3 {
				curStr = v[len(v)-2 : len(v)-1]
			} else {
				strBuilder.WriteString(string(v[len(v)-1]))
				continue
			}
		} else if len(v) == 1 {
			strBuilder.WriteString(v)
			continue
		} else {
			curStr = string(v[0])
		}
		curFactor, _ := strconv.Atoi(string(v[len(v)-1]))
		if curFactor == 0 {
			continue
		}
		preparedStr := strings.Repeat(curStr, curFactor)
		strBuilder.WriteString(preparedStr)
	}

	return strBuilder.String(), nil
}

func regexpValidate(str string) error {
	if str == "" {
		return nil
	}
	if regexp.MustCompile(`(^\d|[^\\]{1}\d\d|\\[[:alpha:]])`).MatchString(str) {
		return ErrInvalidString
	}
	return nil
}

func splitByRegex(str string) []string {
	return regexp.MustCompile(`[[:alpha:]|[:punct:]]{1}\d{0,2}`).FindAllString(str, -1)
}
