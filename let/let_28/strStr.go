package let_28

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if n == m {
		if haystack == needle {
			return 0
		}
		return -1
	} else if n < m {
		return -1
	}

	i := 0
	j := 0

	for i < n && j < m {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == m {
		return i - j
	} else {
		return -1
	}
}

func strStr_2(haystack string, needle string) int {
	return 0
}
