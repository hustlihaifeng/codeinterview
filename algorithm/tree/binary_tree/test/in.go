package main

import (
	"algorithm/tree/binary_tree"
)

var commonBT bt.BinaryTree

func init() {
	GetCommonBT(&commonBT)
}

func GetCommonBT(root *bt.BinaryTree) {
	*root = bt.BinaryTree{Val: 1}
	left := &bt.BinaryTree{Val: 2}
	right := &bt.BinaryTree{Val: 3}
	root.Left = left
	root.Right = right
	left.Right = &bt.BinaryTree{Val: 4}
}
