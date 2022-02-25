package main

import (
	"Letcode_GO/let_1143"
	"fmt"
)

func main() {
	text1 := "pmjghexybyrgzczy"
	text2 := "hafcdqbgncrcbihkd"
	res := let_1143.LongestCommonSubsequence(text1, text2)
	res2 := let_1143.LongestCommonSubsequenceByDp(text1, text2)
	fmt.Println(res)
	fmt.Println(res2)
}
