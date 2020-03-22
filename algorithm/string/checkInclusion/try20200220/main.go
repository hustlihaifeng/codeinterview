package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Printf("s1:%v s2:%v %v\n", s1, s2, checkInclusion(s1, s2))
}

/*
思路：使用桶排序，计算字符个数，只要第一个字符的桶排序和第二个字符的相同长度的串的桶排序相同，第二个的该字串就是第一个的一个排列
伪代码：
第一个长则False
得到第一个的桶排序map1
得到第二个前缀桶排序map2
if mapEqual(map1,map2){
    return true
}
for(i:=len(s1);i<len(s2);i++){
    bgnIdx := i- len(s1)
    map2[bgnIdx]--
    map1[i]++
    if mapEqual(map1,map2){
        return true
    }
}
return false
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	map1 := getMap(s1)
	map2 := getMap(s2[:len(s1)])
	if mapEqual(map1, map2) {
		fmt.Printf("s1 == s2\n")
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		bgnIdx := i - len(s1)
		map2[s2[bgnIdx]]--
		map2[s2[i]]++
		if mapEqual(map1, map2) {
			fmt.Printf("i:%v %v\n", i, s2[bgnIdx+1:i+1])
			return true
		}
	}
	return false
}

func mapEqual(m1, m2 map[byte]int) bool {
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func getMap(s string) map[byte]int {
	rst := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		rst[ch]++
	}
	return rst
}
