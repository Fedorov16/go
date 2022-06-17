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

	splitByRegexp := SplitByRegex(str)

	var strBuilder strings.Builder
	var curStr string
	for _, v := range splitByRegexp {
		if string(v[0]) == "\\" {
			if len(v) == 3 {
				curStr = v[len(v)-2 : len(v)-1]
				curFactor, _ := strconv.Atoi(string(v[len(v)-1]))
				if curFactor != 0 {
					preparedStr := strings.Repeat(curStr, curFactor)
					strBuilder.WriteString(preparedStr)
				}
				continue
			}
			strBuilder.WriteString(string(v[len(v)-1]))
			continue
		} else if len(v) == 1 {
			strBuilder.WriteString(v)
			continue
		}

		curStr = string(v[0])
		curFactor, _ := strconv.Atoi(string(v[len(v)-1]))
		if curFactor != 0 {
			preparedStr := strings.Repeat(curStr, curFactor)
			strBuilder.WriteString(preparedStr)
		}
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

func SplitByRegex(str string) []string {
	return regexp.MustCompile(`[[:alpha:]|[:punct:]]{1}\d{0,2}`).FindAllString(str, -1)
}
