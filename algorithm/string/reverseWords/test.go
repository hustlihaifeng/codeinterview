package reverseWords

import (
	"algorithm/string/string_test"
	"strings"
)

func TestReverseWords() {
	inSli := []string{
		"the sky is blue",
		"  hello world!  ",
		"a good   example",
		"",
	}
	outSli := []string{
		"blue is sky the",
		"world! hello",
		"example good a",
		"",
	}
	string_test.OneInOneOutTest(inSli, outSli, strings.EqualFold, reverseWords)
}
