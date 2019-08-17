package reverseWords

import (
	"strings"
)

/*
# 思路：
1. split by 空格
2. 去除空白串
2. 翻转
3. 输出
# 伪代码
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
*/
func reverseWords(s string) string {
	// 注意：Split的分隔符只能指定一个
	// strSli := strings.Split(s, " \t\r\n")
	// 按空白字符分割：`func Fields(s string) []string`  按指定规则分割：`func FieldsFunc(s string, f func(rune) bool) []string`
	strSli := strings.Fields(s)
	lidx := 0
	for ridx := 0; ridx < len(strSli); ridx++ {
		// 去除首尾空白字符：`func TrimSpace(s string) string`  指定符号集合：`func Trim(s string, cutset string) string`   指定匹配函数：`func TrimFunc(s string, f func(rune) bool) string`
		item := strings.TrimSpace(strSli[ridx])
		if len(item) > 0 {
			strSli[lidx] = item
			lidx++
		}
	}
	strSli = strSli[:lidx]

	maxIdx := len(strSli) - 1
	for i := 0; i < (maxIdx+1)/2; i++ {
		swap(strSli, i, maxIdx-i)
	}
	rst := strings.Join(strSli, " ")
	return rst
}

func swap(strSli []string, l, r int) {
	tmp := strSli[l]
	strSli[l] = strSli[r]
	strSli[r] = tmp
}
