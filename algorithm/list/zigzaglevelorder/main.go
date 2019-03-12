/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (40.22%)
 * Total Accepted:    198.2K
 * Total Submissions: 488.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its zigzag level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
从左到右或从右到左交替变换的将list里面的元素加入rst
从左到右pop，先左后右将子节点加入新list，用新list替代老list。**那么可以保证，list里面的，永远都是从左到右的顺序。**

```go
新建一个list
非空root加入list
for l.Len()>0{
    新建一个sli
    新建一个list
    if l2r {
        从左到右遍历list，将节点加入sli
        先左后右将非空子节点加入新list（PushBack，下一层是从左到右的方式）
        l2r=false
    }else{ // 链表访问方向必须和子节点添加方向一致，然后反向访问时后进先出就可以了（PushFront）
        从右到左遍历，将节点加入sli
        先右后左将非空子节点pushFront到新list（此时新list中，下一层的节点，还是从左到右的顺序）
        l2r=true
    }
    将rst加入rst
    用新list替换老list
}
```
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	rst := [][]int{}

	l := list.New()
	l2r := true
	pushBackIfNotNil(l, root)
	for l.Len() > 0 {
		l2 := list.New()
		var sli []int
		if l2r {
			for l.Len() > 0 {
				pnode := l.Remove(l.Front()).(*TreeNode)
				sli = append(sli, pnode.Val)
				pushBackIfNotNil(l2, pnode.Left)
				pushBackIfNotNil(l2, pnode.Right)
			}
		} else {
			for l.Len() > 0 {
				pnode := l.Remove(l.Back()).(*TreeNode)
				sli = append(sli, pnode.Val)
				pushFrontIfNotNil(l2, pnode.Right)
				pushFrontIfNotNil(l2, pnode.Left)
			}
		}
		l2r = !l2r
		rst = append(rst, sli)
		l = l2
	}

	return rst
}

func pushBackIfNotNil(l *list.List, node *TreeNode) {
	if node != nil {
		l.PushBack(node)
	}
}
func pushFrontIfNotNil(l *list.List, node *TreeNode) {
	if node != nil {
		l.PushFront(node)
	}
}
