package let_offer_58

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	if s == "" {
		return s
	}

	strSlice := strings.Split(s, " ")
	l := len(strSlice)

	var tmpStrs []string
	for i := l - 1; i >= 0; i-- {
		if strSlice[i] == "" {
			continue
		}
		tmpStrs = append(tmpStrs, strSlice[i])
	}

	return strings.Join(tmpStrs, " ")
}
