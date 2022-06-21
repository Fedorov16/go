package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "fr-3fr", expected: "fr---fr"},
		{input: "qwe3", expected: "qweee"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCheckRegexpSplitString(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{input: "a4bc2d5e", expected: []string{"a4", "b", "c2", "d5", "e"}},
		{input: "abccd", expected: []string{"a", "b", "c", "c", "d"}},
		{input: "fr-3fr", expected: []string{"f", "r", "-3", "f", "r"}},
		{input: "qwe3", expected: []string{"q", "w", "e3"}},
		{input: "aaa0b", expected: []string{"a", "a", "a0", "b"}},
		{input: `qwe\4\5`, expected: []string{"q", "w", "e", "\\4", "\\5"}},
		{input: `qwe\45`, expected: []string{"q", "w", "e", "\\45"}},
	}
	for _, stCase := range tests {
		stCase := stCase
		t.Run(stCase.input, func(t *testing.T) {
			result := SplitByRegex(stCase.input)
			require.Equal(t, stCase.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "0", "--23--", "_0-9=43", "\\e", "q3\\we"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
