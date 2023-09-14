package let_5

func longestPalindrome(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}

	ans := string(s[0])
	for i := 0; i < l; i++ {
		j := i + 2
		for j <= l {
			tmp := s[i:j]
			if IsPalindrome(tmp) {
				ans = max(ans, tmp)
			}
			j++
		}
	}
	return ans
}

func IsPalindrome(s string) bool {
	l := len(s)
	if l == 1 {
		return true
	}

	left := 0
	right := l - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
