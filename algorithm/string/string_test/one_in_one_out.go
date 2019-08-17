package string_test

import "fmt"

func OneInOneOutTest(inSli, outSli []string,
	compareFunc func(a, b string) bool,
	targetFunc func(string) string) {

	Assert(len(inSli) == len(outSli), "输入输出个数不一样")

	for idx, in := range inSli {
		out := outSli[idx]
		realOut := targetFunc(in)
		if compareFunc(out, realOut) {
			fmt.Printf("  pass:%v\n", in)
		} else {
			fmt.Printf("  fail:%v\nexpect:%v\n  real:%v\n", in, out, realOut)
		}
	}
}
func Assert(cond bool, err_msg string) {
	if !cond {
		panic(err_msg)
	}
}
