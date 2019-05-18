package main

func main() {

}

/*思路：
用map记录每个字符出现的索引，对一个字符，如果其上次出现的索引在本次查询的开始位置前，那么出现重复，更新开始位置为上次出现的位置加一，继续找；否则不重复，更新当前长度，并尝试更最大值。
*/
/*伪代码：
bgnIdx=0
maxLen=0
for i处的字符ch{
    if ch在map中出现 && oldi >= bgnIdx {// 重复
        bgnIdx = oldi+1
    }else{ // 1. ch在map中未出现 2. ch出现单在本次bgnIdx之前。都不会导致本次重复
        当前长度为i-bgnIdx+1
        if 当前长度>maxLne {
            更新maxLen
        }
	}
	更新map中ch的索引为i
}
*/
func lengthOfLongestSubstring(s string) int {
	bgnIdx := 0
	maxLen := 0
	idxMap := make(map[rune]int)
	for i, ch := range s {
		if oldi, ok := idxMap[ch]; ok && oldi >= bgnIdx {
			bgnIdx = oldi + 1
		} else {
			cLen := i - bgnIdx + 1
			if cLen > maxLen {
				maxLen = cLen
			}
		}
		idxMap[ch] = i
	}
	return maxLen
}
