package bt

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func TreeInvalid(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	return false
}
