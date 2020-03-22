package main

import "container/list"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
# 问题
1. 二叉树的锯齿形遍历，顺序交替替换

# 思路
1. 先深度优先，不同省的存入不同的数组，然后索引为基数的数组进反转
2. 遍历复杂度至少为O(n),反转复杂度为O(n)

*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	// TODO:
	init := dfs(root)
	return reverseOdd(init)
}

type dfsNode struct {
	h int
	n *TreeNode
}

func dfs(root *TreeNode) (rst [][]int) {
	// 因为要计算深度，所以是先访问然后加入fifo队列
	lst := list.New()
	// TODO:deal root
	visit(root, 0, &rst, lst)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Front()).(dfsNode)
		// 判断是否是nil，处理深度，尝试新建并加入对应深度的lst,加入队列
		visit(cnode.n.Left, cnode.h+1, &rst, lst)
		visit(cnode.n.Right, cnode.h+1, &rst, lst)
	}
	return rst
}

// 判断是否是nil，处理深度，尝试新建并加入对应深度的lst,加入队列
func visit(root *TreeNode, h int, rst *[][]int, lst *list.List) {
	if root == nil {
		return
	}
	if h == len(*rst) {
		*rst = append(*rst, []int{root.Val})
	} else {
		(*rst)[h] = append((*rst)[h], root.Val)
	}
	dn := dfsNode{h, root}
	lst.PushBack(dn)
}

func reverseOdd(in [][]int) (rst [][]int) {
	for idx, sli := range in {
		if idx%2 == 1 {
			in[idx] = reverseSli(sli)
		}
	}
	return in
}

func reverseSli(in []int) (rst []int) {
	for lidx, ridx := 0, len(in)-1; lidx < ridx; {
		in[lidx], in[ridx] = in[ridx], in[lidx]
		lidx++
		ridx--
	}
	return in
}
