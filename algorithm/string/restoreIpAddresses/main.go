package main

import (
	"strconv"
	"strings"
)

func main() {

}

/*
思路：运用递归，从左到右，合并1 2 3 个数字的情况，如果子调用返回的不是空的话。
伪代码：
func restoreIpAddressesRecursion(s string,d int) [][]string {
	if d == 1 {
		if isIPItem(s) {
			return [][]string{[]string{s}}
		}
		return nil
	}
	var rst [][]string
	if len(s) > 0 {
		item := s[:1]
		if isIPItem(item){
			chdSli := restoreIpAddressesRecursion(s[1:],d-1)
			if chdSli != nil {
				// 在每一个以为数组中最前面加上item
				// rst 中加上chdSli
			}
		}
	}
	if len(s) > 1 {
		item := s[:2]
		if isIPItem(item){
			chdSli := restoreIpAddressesRecursion(s[2:],d-1)
			if chdSli != nil {
				// 在每一个以为数组中最前面加上item
				// rst 中加上chdSli
			}
		}
	}
	if len(s) > 2 {
		item := s[:3]
		if isIPItem(item){
			chdSli := restoreIpAddressesRecursion(s[3:],d-1)
			if chdSli != nil {
				// 在每一个以为数组中最前面加上item
				// rst 中加上chdSli
			}
		}
	}
	return rst
}
*/
func restoreIpAddresses(s string) []string {
	sli2 := restoreIpAddressesRecursion(s, 4)
	var rst []string
	for _, sli := range sli2 {
		sli = reverseSli(sli)
		rst = append(rst, strings.Join(sli, "."))
	}
	return rst
}

func reverseSli(sli []string) (rst []string) {
	rst = make([]string, len(sli))
	for idx, item := range sli {
		rst[len(sli)-1-idx] = item
	}
	return rst
}

func isIPItem(s string) bool {
	if s == "0" {
		return true
	} else if strings.HasPrefix(s, "0") {
		return false
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if 0 <= num && num <= 255 {
		return true
	}
	return false
}
func addItemToEverySli(s string, sli2 [][]string) {
	for idx, sli := range sli2 {
		sli2[idx] = append(sli, s)
	}
}
func restoreIpAddressesRecursion(s string, d int) [][]string {
	if d == 1 {
		if isIPItem(s) {
			sli := make([]string, 0, 4)
			sli = append(sli, s)
			return [][]string{sli}
		}
		return nil
	}

	var rst [][]string
	for i := 1; i <= 3; i++ {
		if len(s) >= i {
			item := s[:i]
			if isIPItem(item) {
				chdSli := restoreIpAddressesRecursion(s[i:], d-1)
				if chdSli != nil {
					addItemToEverySli(item, chdSli)
					rst = append(rst, chdSli...)
				}
			}
		}
	}
	return rst
}
