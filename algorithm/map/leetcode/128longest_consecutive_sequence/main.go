package main

import "fmt"

func main() {
	var nums []int
	var output int

	nums = []int{100, 4, 200, 1, 3, 2}
	output = 4
	fmt.Printf("%v %v expect %v\n", nums, longestConsecutive(nums), output)
}

/*
hash map读取都是O(1)的操作，那么可以把切片中元素都唯一的加入到map中，然后取出元素。从数组中找其相邻元素，没找到则不连续。每个连续块会有首尾两个空位，但总体次数不超过元素个数的三倍。所以是O(n)
*/
/*
for 数组里面的每一个元素{
    覆盖式的加入map，值设为true
}

最大计数设置为0
for map中的每一对kv元祖 {
    计数设为0

    计数加一并从map中删除k
    khigh,klow设为k
    for ++khigh在map中存在 {
        计数加一并从map中删除khigh
    }

    for --klow在map中存在 {
        计数加一并从map中删除klow
    }

    if 计数大于最大计数{
        最大计数设为当前计数
    }
}
*/
func longestConsecutive(nums []int) int {
	if longestConsecutiveInputInvalid(nums) {
		return 0
	}

	hmap := make(map[int]bool)
	for _, num := range nums {
		hmap[num] = true
	}

	maxCnt := 0
	for k, _ := range hmap {
		cnt := 0
		incrCntAndDelKey(&cnt, hmap, k)

		khigh, klow := k, k
		for true {
			khigh++
			if _, ok := hmap[khigh]; !ok {
				break
			}
			incrCntAndDelKey(&cnt, hmap, khigh)
		}

		for true {
			klow--
			if _, ok := hmap[klow]; !ok {
				break
			}
			incrCntAndDelKey(&cnt, hmap, klow)
		}

		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	return maxCnt
}

func longestConsecutiveInputInvalid(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}

	return false
}

func incrCntAndDelKey(pcnt *int, hmap map[int]bool, key int) {
	(*pcnt)++
	delete(hmap, key)
}
