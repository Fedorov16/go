package hw03frequencyanalysis

import (
	"strings"
)

type SD []int
type SW []string

func (s *SD) Add(sStr *SW, newVal string, newInt int) {
	if len(*s) == 0 {
		*s = append(*s, newInt)
		*sStr = append(*sStr, newVal)
		return
	}
	for i, val := range *s {
		oldS := *s
		oldStr := *sStr
		if newInt > val {
			// ints
			partBefore := make([]int, len(oldS[:i]))
			partAfter := make([]int, len(oldS[i:]))
			copy(partBefore, oldS[:i])
			copy(partAfter, oldS[i:])
			r := append(partBefore, newInt)
			*s = append(r, partAfter...)

			// strings
			partStringBefore := make([]string, len(oldStr[:i]))
			partStringAfter := make([]string, len(oldStr[i:]))
			copy(partStringBefore, oldStr[:i])
			copy(partStringAfter, oldStr[i:])
			rStr := append(partStringBefore, newVal)
			*sStr = append(rStr, partStringAfter...)

			return
		}
	}
	*s = append(*s, newInt)
	*sStr = append(*sStr, newVal)
}

func Top10(text string) []string {
	matches := strings.Fields(text)
	m := make(map[string]int)
	for _, v := range matches {
		m[v]++
	}
	sInt := new(SD)
	sStr := new(SW)
	for str, count := range m {
		sInt.Add(sStr, str, count)
	}
	res := *sStr
	if len(res) == 0 {
		return nil
	}

	return res[:10]
}
