# 问题
1. 两个树节点的最近祖先

# 思路
1. 如果保存了父节点，那么就是找两个链表的交叉点问题

2. 没有保存父节点：递归找，lowestCommonAncestor返回子树中是否至少找到一个点，如果左右子树都找到，那么就是当前点；如果只有一个子树找到，要么是都找到了，要么是只找到一个，返回当前点；如果都没找到，那么当前树没有。

3. 不是很有道理，把这算法给记着把。不过树的题目，一般都是递归

# 伪代码：
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	lrst := lowestCommonAncestor(root.Left, p, q)
	rrst := lowestCommonAncestor(root.Right, p, q)
	if lrst != nil && rrst != nil {
		return root
	}
	if lrst != nil {
		return lrst
	}
	return rrst
}
```