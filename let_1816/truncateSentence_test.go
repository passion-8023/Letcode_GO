package let_1816

import "testing"

func TestTruncateSentence(t *testing.T) {
	str := "Hello how are you Contestant"
	k := 2
	newStr := TruncateSentence3(str, k)
	t.Logf(newStr)
}
