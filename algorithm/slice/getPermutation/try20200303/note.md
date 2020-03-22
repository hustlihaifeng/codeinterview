# 问题
找[1...n]的第k个排列：

# 思路
找打最大的m，使得M!<=k,那么k/k!是确定第一行的第几大的数。接着进行递归，k=k%m!,sortSli中去除上一次的元素。注意k==m!的时候怎么处理

# 伪代码
```go
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
```

# 详见
1. [main.go](main.go)