package main

import "fmt"

func main() {
	sliceFromArray()
	fmt.Println()
	sliceFromSlice()
}

func sliceFromArray() {
	ary := [4]int{10, 20, 30, 40}
	sli := ary[0:2]
	newSli := append(sli, 50)
	fmt.Println(ary)
	fmt.Println(sli)
	fmt.Println(newSli)
}

func sliceFromSlice() {
	ary := []int{10, 20, 30, 40}
	sli := ary[0:2]
	newSli := append(sli, 50)
	fmt.Println(ary)
	fmt.Println(sli)
	fmt.Println(newSli)
}

/*
结果：
[10 20 50 40]
[10 20]
[10 20 50]

[10 20 50 40]
[10 20]
[10 20 50]

*/
