package main

import "fmt"

func main() {
	num2 := "456"
	num1 := "123"
	out := "56088"
	fmt.Printf("num1:%v num2:%v out:%v real:%v\n", num1, num2, out, multiply(num1, num2))
}

/*
思路：运用竖式乘法，计算mXn次，结果依次累加，索引和即每次结果后面改加的0的个数。
伪代码：
rst := "0"
for idx1:=0;idx1<len(num1);idx1++ {
	offset := 0
	for idx2:=0;idx2<len(num2);idx2++{
		n1=num1[idx1]-'0'
		n2=num2[idx2]-'0'
		sum := n1*n2+offset
		offset = sum/10
		left := sum%10
		rst = add(rst,left,idx1+idx2)
	}
	rst = add(rst,offset,idx1+len(num2))
}

func add(big string, small, zeronum int) string {
	if len(big)<(zeronum+1){
		big前面补0到zeronum+1位数
	}
	big前面补一个0，来处理最后的进位
	offset := small
	for idx:=zeronum;idx<len(big);idx++{
		n := big[idx]-'0'
		sum = n + offset
		offset = sum/10
		big[idx]=sum%10+'0'
	}
	去除big前面的0
	返回big
}
*/

func multiply(num1 string, num2 string) string {
	rst := "0"
	for idx1 := len(num1) - 1; idx1 >= 0; idx1-- {
		offset := 0
		for idx2 := len(num2) - 1; idx2 >= 0; idx2-- {
			n1 := num1[idx1] - '0'
			n2 := num2[idx2] - '0'
			sum := int(n1)*int(n2) + offset
			offset = sum / 10
			left := sum % 10
			rst = add(rst, left, len(num1)-1+len(num2)-1-idx1-idx2)
		}
		rst = add(rst, offset, len(num1)-1+len(num2)-idx1)
	}
	rst = removeLeftZero(rst)
	if len(rst) == 0 {
		return "0"
	}
	return rst
}

func Assert(con bool, msg string) {
	if !con {
		panic(msg)
	}
}
func add(big string, small, zeronum int) (r string) {
	// Assert(zeronum >= 0, "zeronum>=0")
	// defer func() {
	// 	fmt.Printf("%v+%v*10^%v=%v\n", big, small, zeronum, r)
	// }()
	var rst []byte
	if len(big) < (zeronum + 1) {
		// big前面补0到zeronum+1位数
		rst = addLeftZero(big, zeronum+1-len(big))
	} else {
		rst = []byte(big)
	}
	// big前面补一个0，来处理最后的进位
	rst = addLeftZero(string(rst), 1)
	offset := small
	for idx := len(rst) - zeronum - 1; idx >= 0; idx-- {
		n := rst[idx] - '0'
		sum := int(n) + offset
		offset = sum / 10
		rst[idx] = byte(sum%10) + '0'
	}
	r = removeLeftZero(string(rst))
	if len(r) == 0 {
		return "0"
	}
	return r
}

func removeLeftZero(num string) (rst string) {
	for idx := 0; idx < len(num); idx++ {
		if num[idx] != '0' {
			return string(num[idx:])
		}
	}
	return "0"
}
func addLeftZero(num string, zeronum int) (rst []byte) {
	// defer func() {
	// 	fmt.Printf("addLeftZero: num:%v zeronum:%v out:%v\n", num, zeronum, string(rst))
	// }()
	rst = make([]byte, len(num)+zeronum)
	for i := 0; i < zeronum; i++ {
		rst[i] = '0'
	}
	for i := 0; i < len(num); i++ {
		rst[zeronum+i] = num[i]
	}
	return rst
}
