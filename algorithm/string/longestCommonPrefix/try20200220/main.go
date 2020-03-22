package main

import "fmt"

func main() {
	in := []string{"flower", "flow", "flight"}
	fmt.Printf("in:%v out:%v\n", in, longestCommonPrefix(in))
}

/*
思路：
遍历字符串数组，更新当前位置索引，对比各个字符串当前位置值是否相等
伪代码：
idx:=0
for true {
	for i,str := range strs {
		if len(str) <= idx {
			return str[:idx]
		}
		if i == 0 {
			ch = str[i]
		}else{
			if str[i]!=ch {
				return str[:idx]
			}
		}
	}
	idx++
}
return ""

*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	idx := 0
	for true {
		var ch byte
		for i, str := range strs {
			if len(str) <= idx {
				fmt.Printf("len: idx:%v str:%v i:%v ch:%v\n", idx, str, i, string([]byte{ch}))
				return str[:idx]
			}
			if i == 0 {
				ch = str[idx]
			} else {
				if str[idx] != ch {
					fmt.Printf("diff: idx:%v str:%v i:%v ch:%v\n", idx, str, i, string([]byte{ch}))
					return str[:idx]
				}
			}
		}
		idx++
	}
	return ""
}
