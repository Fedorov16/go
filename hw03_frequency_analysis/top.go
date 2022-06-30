package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type CustomMap struct {
	word  string
	count int
}

const maxRes = 10

func Top10(text string) []string {
	matches := strings.Fields(text)
	r := regexp.MustCompile(`[.,!']`)
	m := make(map[string]int)
	for _, v := range matches {
		str := r.ReplaceAllString(v, "")
		if str == "-" {
			continue
		}
		m[strings.ToLower(str)]++
	}
	sliceCM := make([]CustomMap, 0, len(m))
	for i, v := range m {
		sliceCM = append(sliceCM, CustomMap{i, v})
	}

	sort.Slice(sliceCM, func(i, j int) bool {
		if sliceCM[i].count == sliceCM[j].count {
			return sliceCM[i].word < sliceCM[j].word
		}
		return sliceCM[i].count > sliceCM[j].count
	})

	if len(sliceCM) == 0 {
		return nil
	}

	maxVal := maxRes
	if maxRes > len(m) {
		maxVal = len(m)
	}

	res := make([]string, maxVal, len(m))
	for i, v := range sliceCM[:maxVal] {
		res[i] = v.word
	}

	return res
}
