package hw03frequencyanalysis

import (
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
	m := make(map[string]int)
	for _, v := range matches {
		m[v]++
	}
	sliceCM := make([]CustomMap, 0)
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

	res := make([]string, maxRes)
	for i, v := range sliceCM[:maxRes] {
		res[i] = v.word
	}

	return res
}
