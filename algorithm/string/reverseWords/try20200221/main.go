package main

import "strings"

func main() {

}

/*
思路：
1. 按空格分隔得到字符串数组
2. 去除空串
3. 翻转数组
4. join
伪代码：
strSli:= string.Split(s," ")
rst := make([]string,len(strSli),0)
for i:= len(strSli); i>=0;i-- {
	if str != "" {
		rst = append(rst,str)
	}
}
return strings.Join(rst," ")
*/

func reverseWords(s string) string {
	strSli := strings.Split(s, " ")
	rst := make([]string, 0, len(strSli))
	for i := len(strSli) - 1; i >= 0; i-- {
		str := strSli[i]
		if str != "" {
			rst = append(rst, str)
		}
	}
	return strings.Join(rst, " ")
}
