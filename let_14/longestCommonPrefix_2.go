package let_14

func longestCommonPrefix_2(strs []string) string {
	if len(strs) <= 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	result := make([]byte, 0)
	for i := 0; i < minLen; i++ {
		ch := strs[0][i]
		for _, str := range strs[1:] {
			if ch != str[i] {
				return string(result)
			}
		}
		result = append(result, ch)
	}
	return string(result)
}
