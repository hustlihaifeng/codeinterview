package main

import "fmt"

func main() {

}

/*
# 问题
找[1...n]的第k个排列：

# 思路
找打最大的m，使得M!<=k,那么k/k!是确定第一行的第几大的数。接着进行递归，k=k%m!,sortSli中去除上一次的元素。注意k==m!的时候怎么处理

# 伪代码
sli := [1...n]
m,mFactorial := findMaxFactorial(n,k)
remainder := k%mFactorial
quotient := k/mFactorial
curNum := sli[quotient-1]
return curNum+recursive(sli,remainder)

func recursive(sli []int,k int) (rst string){
	if len(sli)==1 {
		return sli[0]
	}
	mFactorial := Factorial(len(sli)-1)
	remainder := k%mFactorial
	quotient := k/mFactorial
	if remainder == 0 {
		remainder=mFactorial
		quotient--
	}
	curNum := sli[quotient]
	newsli := append(sli[:quotient])
	if quotient < len(sli)-1 {
		newsli = append(newsli,sli[quotient+1:]...)
	}
	return curNum+recursive(newsli,remainder)
}

*/

func getPermutation(n int, k int) string {
	sli := make([]string, n, n)
	for idx := range sli {
		sli[idx] = fmt.Sprintf("%v", idx+1)
	}

	factorial := make([]int, n-1, n-1)
	old := 1
	for i := 1; i <= n-1; i++ {
		old = old * i
		factorial[i-1] = old
	}

	return recursive(sli, factorial, k)
}

func recursive(sli []string, factorial []int, k int) (rst string) {
	if len(sli) == 1 {
		return sli[0]
	}
	mFactorial := factorial[len(sli)-2]
	remainder := k % mFactorial
	quotient := k / mFactorial
	if remainder == 0 { // 前面有 quotient个n-1排列，最后一个里面的第 remainder 个
		remainder = mFactorial
		quotient--
	}
	curNum := sli[quotient]
	newsli := append(sli[:quotient])
	if quotient < len(sli)-1 { // 最后一个的话，quotient+1超过数组最大下标
		newsli = append(newsli, sli[quotient+1:]...)
	}
	return curNum + recursive(newsli, factorial, remainder)
}
