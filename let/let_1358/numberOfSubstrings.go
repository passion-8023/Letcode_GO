package let_1358

// 方法一
// 算法正确，但是超出时间限制
func numberOfSubstrings(s string) int {
	var count int
	start := 0
	end := start + 2
	for start <= len(s)-3 {
		subString := s[start : end+1] // 获取子串

		tmpMap := make(map[string]struct{})
		for _, item := range subString {
			tmpMap[string(item)] = struct{}{}
		}
		if len(tmpMap) == 3 {
			count++
		}

		if end == len(s)-1 {
			if end-start == 2 {
				// 结束
				break
			} else {
				start++
				end = start + 2
			}
		} else {
			end++
		}
	}

	return count
}

// 方法二
// 算法正确，但是超出时间限制
func numberOfSubstrings2(s string) int {
	var (
		count int
		start int
		end   int
	)

	tmpMap := make(map[string]struct{})

	for start <= len(s)-3 {
		tmpMap[string(s[end])] = struct{}{}
		if len(tmpMap) == 3 {
			count++
		}
		if end == len(s)-1 {
			if end-start == 2 {
				// 结束
				break
			} else {
				start++
				end = start
				tmpMap = make(map[string]struct{})
			}
		} else {
			end++
		}
	}

	return count
}

// 方法三
// 算法正确，但是超出时间限制
func numberOfSubstrings3(s string) int {
	var (
		count int
		start int
		end   int
	)

	tmpMap := make(map[string]struct{})

	for start < len(s)-2 && end < len(s) {
		tmpMap[string(s[end])] = struct{}{}
		if len(tmpMap) == 3 {
			count += 1 + len(s) - end - 1
			start++
			end = start
			tmpMap = make(map[string]struct{})
		} else {
			end++
		}
	}

	return count
}

// 方法四
func numberOfSubstrings4(s string) int {
	var (
		count int
		start int
		end   int
	)

	end = 3

	tmpMap := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
	}

	for i := 0; i < end; i++ {
		tmpMap[string(s[i])]++
	}

	for start < len(s)-2 {
		if tmpMap["a"] != 0 && tmpMap["b"] != 0 && tmpMap["c"] != 0 {
			count += len(s) - end + 1
			tmpMap[string(s[start])]--
			start++
		} else {
			end++
			if end > len(s) {
				break
			}
			tmpMap[string(s[end-1])]++
		}
	}

	return count
}

// 方法五
// 大神提交的代码 运行时间0ms
func numberOfSubstrings5(s string) int {
	nos := 0
	last := []int{-1, -1, -1}
	cnt, j := 0, -1
	for i := 0; i < len(s); i++ {
		cur := int(s[i] - 'a')
		if last[cur] == -1 {
			cnt++
		}
		last[cur] = i
		if cnt == 3 {
			// bbabc
			nos += min(last[(cur+1)%3], last[(cur+2)%3]) + 1
			j = i + 1
			break
		}
	}
	if j == -1 { // 不存在满足条件的子字符串
		return 0
	}
	for i := j; i < len(s); i++ {
		cur := int(s[i] - 'a')
		last[cur] = i

		leftmost := i
		for _, p := range last {
			leftmost = min(leftmost, p)
		}
		nos += leftmost + 1
	}
	return nos
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
