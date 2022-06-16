package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if !regexpValidate(str) {
		return "", ErrInvalidString
	}

	splitByRegexp := splitByRegex(str)

	var strBuilder strings.Builder
	var curStr string
	for i, v := range splitByRegexp {
		if len(v) == 1 {
			if string(v[0]) == "\\" && splitByRegexp[i-1] == "\\" {
				strBuilder.WriteString(v)
			} else if string(v[0]) == "\\" {
				continue
			}
			strBuilder.WriteString(v)
			continue
		}
		if string(v[0]) == "\\" {
			if len(v) == 3 {
				curStr = v[len(v)-2 : len(v)-1]
			} else if string(splitByRegexp[i-1]) == "\\" {
				curStr = "\\"
			} else {
				strBuilder.WriteString(string(v[len(v)-1]))
				continue
			}
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

func regexpValidate(str string) bool {
	return !regexp.MustCompile(`(^\d|[^\\]{1}\d\d|\\[[:alpha:]])`).MatchString(str)
}

func splitByRegex(str string) []string {
	return regexp.MustCompile(`[[:alpha:]|[:punct:]]{1}\d{0,2}`).FindAllString(str, -1)
}
