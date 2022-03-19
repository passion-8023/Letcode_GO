package let_1816

import (
	"strings"
)

func TruncateSentence(s string, k int) string {
	strSlice := strings.Split(s, " ")
	//fmt.Printf("%q (nil = %v)\n", strSlice, strSlice == nil)
	strSlice = strSlice[:k]
	return strings.Join(strSlice, " ")
}

func TruncateSentence2(s string, k int) string {
	var n int
	for i, val := range s {
		//fmt.Printf("%d-%s\n", i, string(val))
		if string(val) == " " {
			n++
		}
		if n == k {
			k = i
			break
		}
		if i == len(s)-1 {
			k = len(s)
			break
		}
	}
	return s[:k]
}

func TruncateSentence3(s string, k int) string {
	var i int
	for ; i < len(s); i++ {
		if rune(s[i]) == ' ' {
			k--
		}
		if k == 0 {
			break
		}
	}
	return s[:i]
}

