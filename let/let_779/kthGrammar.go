package let_779

//func kthGrammar(n int, k int) int {
//	str := string(callback(n)[k-1])
//
//	result, _ := strconv.Atoi(str)
//	return result
//}
//
//func callback(n int) string {
//	if n == 1 {
//		return "0"
//	}
//
//	if n == 2 {
//		return "01"
//	}
//
//	var result string
//	for _, item := range callback(n - 1) {
//		if item == '0' {
//			result += "01"
//		}
//		if item == '1' {
//			result += "10"
//		}
//	}
//
//	return result
//}

// 分析：
// 第1行 0	= 2^0
// 第2行 01	= 2^1 增加了 2^0 个数
// 第3行 0110	= 2^2 增加了 2^1 个数
// 第4行 0110 1001	= 2^3 增加了 2^2 个数
// 第5行 01101001 10010110	= 2^4 增加了 2^3 个数
func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}

	if k > 1<<(n-2) {
		// 值在后半部分
		// 对前半部分取反就是后半部分
		return 1 ^ kthGrammar(n-1, k-1<<(n-2))
	}

	return kthGrammar(n-1, k)
}
