package main

import "fmt"

func main() {
	testCopy()
}

func testCopy() {
	src := []int{1, 2, 3, 4}

	dst := make([]int, 2, 6)
	dst[0] = 5
	dst[1] = 6

	copy(dst, src)

	fmt.Println(src)
	fmt.Println(dst)
}
