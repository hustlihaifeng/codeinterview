package main

import "fmt"

/*
# 思路
用一个map，保存每个字符上次出现的位置。对一个字符：
1. 如果在map中存在且上次出现的位置大于等于当前起始位置，那么该字符终结当前起始位置的查找。如果本次查找的长度大于当前最大长度，那么更新当前最大长度。
2. 否则，当前字符在当前起始位置之后没有出现。当前字符是安全的。

# 伪代码
beginIndex,maxLen:=0
for idx,ch := range s {
	if !IsSafe(indexMap,beginIndex,ch){
		curLen := idx-beginIdex
		maxLen := max(maxLen,curLen)
		beginIndex=idx
	}
	indexMap[ch]=idx
}
return maxLen

func IsSafe(indexMap,beginIndex,ch) bool {
	if oldIdx,ok := indexMap[ch]; !ok {
		return true
	}
	if oldIdx < beginIndex {
		return true
	}
	return false
}
*/

func main() {
	in := "pwwkew"
	fmt.Printf("in:%v out:%v\n", in, lengthOfLongestSubstring(in))
}

func IsSafe(indexMap map[rune]int, beginIndex int, ch rune) bool {
	if oldIdx, ok := indexMap[ch]; !ok {
		return true
	} else {
		if oldIdx < beginIndex {
			return true
		} else {
			return false
		}
	}
}
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	beginIndex := 0
	indexMap := make(map[rune]int)
	for idx, ch := range s {
		if !IsSafe(indexMap, beginIndex, ch) {
			curLen := idx - beginIndex
			if curLen > maxLen {
				maxLen = curLen
				fmt.Printf("maxlen:%v idx:%v\n", maxLen, idx)
			}
			beginIndex = indexMap[ch] + 1
		}
		indexMap[ch] = idx
	}
	curLen := len(s) - beginIndex
	if curLen > maxLen {
		maxLen = curLen
		fmt.Printf("maxlen:%v end\n", maxLen)
	}
	return maxLen
}
