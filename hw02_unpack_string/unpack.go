package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var prev rune
	hasPrevRune := false
	hasPrevSlash := false

	var strBuilder strings.Builder
	for _, r := range str {
		switch {
		case '0' <= r && r <= '9':
			if !hasPrevRune {
				return "", ErrInvalidString
			}
			if !hasPrevSlash {
				for i := rune(0); i < r-'0'; i++ {
					strBuilder.WriteRune(prev)
				}
				hasPrevRune = false
			} else {
				prev = r
				hasPrevSlash = false
			}

		case r == '\\':
			if hasPrevSlash {
				prev = r
				hasPrevSlash = false
			} else {
				if hasPrevRune {
					strBuilder.WriteRune(prev)
				}
				hasPrevSlash = true
			}
		default:
			if hasPrevRune {
				strBuilder.WriteRune(prev)
			}
			prev = r
			hasPrevRune = true
		}
	}

	if hasPrevRune {
		strBuilder.WriteRune(prev)
	}

	return strBuilder.String(), nil
}
