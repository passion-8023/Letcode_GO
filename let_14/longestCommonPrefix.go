package let_14

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 1 {
		return strs[0]
	}

	current := strs[0]

	for i := 1; i < l-1; i++ {
		current = compare(current, strs[i])
	}

	return current
}

func compare(str1, str2 string) string {
	n := len(str1)
	m := len(str2)

	var result []byte
	for i := 0; i < min(n, m); i++ {
		if str1[i] == str2[i] {
			result = append(result, str1[i])
		} else {
			break
		}
	}

	return string(result)
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}
