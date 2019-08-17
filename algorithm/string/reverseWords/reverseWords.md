# 问题描述

1. 字符串里单词间反序输出，去除空格，详见 <https://leetcode-cn.com/problems/reverse-words-in-a-string/>

# 思路和伪代码
## 思路
1. split by 空格
2. 去除空白串
2. 翻转
3. 输出
## 伪代码
```go
strSli := strings.Split(s," \t\r\n")
lidx,ridx := 0
for ridx:=0;ridx<len(strSli);ridx++{
    item := strSli[ridx].Trim()
    if len(item)>0{
        strSli[lidx]=item
        lidx++
    }
}
maxIdx := len(strSli)-1
for i:=0;i<=maxIdx/2; i++ {
    swap(strSli,i,maxIdx-i)
}
rst:= strings.Join(strSli,' ')
```

# 注意点

1. 注意：Split的分隔符只能指定一个，Trim的分隔符是指定一个集合，`strings.Split(s, " \t\r\n")`是错误的用法
2. 分割：

- 按空白字符分割：`func Fields(s string) []string`
- 按指定规则分割：`func FieldsFunc(s string, f func(rune) bool) []string`

3. Trim：

- 去除首尾空白字符：`func TrimSpace(s string) string`
- 指定符号集合：`func Trim(s string, cutset string) string`
- 指定匹配函数：`func TrimFunc(s string, f func(rune) bool) string`

4. 测试函数：

```go
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
```
