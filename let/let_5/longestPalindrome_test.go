package let_5

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	str := "ac"

	palindrome := longestPalindrome_2(str)
	fmt.Println(palindrome)
}
