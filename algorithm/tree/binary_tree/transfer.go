package bt

import (
	"container/list"
	"errors"
	"fmt"
)

/*
经验：
1. 将异常处理 初值处理 边界异常处理 都通过函数预留，最开始只考虑大部分场景的逻辑（即该函数的核心逻辑）。（概要设计时）可以都用函数代替
2. 把自己的设备带着，这样用自己的设备，顺手一些。
3. 代码不超过25行（14寸的vscode里的一个屏幕），否则一定是划分不清晰。一个良好的代码，如下面的，应该只有一个分支判断
*/
func TransferLeftMiddleRight(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}

	/*
		1. first access(nil in visitMap), push all element to stack
		2. second access(false in visitMap, and previous left has been visited), visit
	*/
	lst := list.New()
	visitMap := make(map[*BinaryTree]bool)
	pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		if NodeIsNilOrVisited(cnode.Left, visitMap) {
			visitAndSetMap(cnode, visitMap)
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode.Right)
		} else {
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode)
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode.Left)
		}
	}

	return nil
}

func TransferLeftMiddleRightBad(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}
	/*
	   1. if Left has not been visited, push current BinaryTree,push Left BinaryTree, continue
	   2. if Left has beet visited, visit current; if Right has not been visit,push
	*/
	visitedMap := make(map[*BinaryTree]bool)
	lst := list.New()

	pushBackAdnUpdateVisited(lst, visitedMap, root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		if cnode.Left == nil {
			transferMiddleRight(cnode, lst, visitedMap)
		} else {
			if visited, ok := visitedMap[cnode.Left]; !ok {
				pushBackAdnUpdateVisited(lst, visitedMap, cnode)
				pushBackAdnUpdateVisited(lst, visitedMap, cnode.Left)
				continue
			} else {
				if visited {
					transferMiddleRight(cnode, lst, visitedMap)
					continue
				} else {
					return errors.New("Left not visited")
				}
			}
		}
	}
	return nil
}

func TransferMiddleLeftRight(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}
	lst := list.New()
	PushBackIfNotNil(lst, root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		visit(cnode)
		PushBackIfNotNil(lst, cnode.Right)
		PushBackIfNotNil(lst, cnode.Left)
	}

	return nil
}

func TransferMiddleLeftRightBad(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}

	/*
		for every node:
			1. visit it; push left is left is not nil; push right if right is not nil
	*/
	lst := list.New()
	lst.PushBack(root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		visit(cnode)
		if cnode.Right != nil {
			lst.PushBack(cnode.Right)
		}
		if cnode.Left != nil {
			lst.PushBack(cnode.Left)
		}

	}

	return nil
}

func TransferLeftRightMiddle(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}

	lst := list.New()
	visitMap := make(map[*BinaryTree]bool)

	pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, root)
	// main
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		if NodeIsNilOrVisited(cnode.Left, visitMap) && NodeIsNilOrVisited(cnode.Right, visitMap) {
			visitAndSetMap(cnode, visitMap)
		} else {
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode)
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode.Right)
			pushBackAdnUpdateVisitedIfNotNil(lst, visitMap, cnode.Left)
		}
	}

	return nil
}
func TransferLeftRightMiddleBad(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}

	/*
		for every node:
			if left and right is nil. or left and right have been visited, then visit it
			else push it, if right is not nil, push it; if left is not nil, push it
	*/

	visitedMap := make(map[*BinaryTree]bool)
	lst := list.New()
	pushBackAdnUpdateVisited(lst, visitedMap, root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Back()).(*BinaryTree)
		if cnode.Left == nil && cnode.Right == nil {
			visitAndSetMap(cnode, visitedMap)
		} else if cnode.Left == nil && cnode.Right != nil {
			rvisited, rok := visitedMap[cnode.Right]
			if rok && rvisited {
				visitAndSetMap(cnode, visitedMap)
			} else {
				pushBackAdnUpdateVisited(lst, visitedMap, cnode)
				pushBackAdnUpdateVisited(lst, visitedMap, cnode.Right)
			}

		} else if cnode.Left != nil && cnode.Right == nil {
			lvisited, lok := visitedMap[cnode.Left]
			if lok && lvisited {
				visitAndSetMap(cnode, visitedMap)
			} else {
				pushBackAdnUpdateVisited(lst, visitedMap, cnode)
				pushBackAdnUpdateVisited(lst, visitedMap, cnode.Left)
			}
		} else { // both is not nil
			lvisited, lok := visitedMap[cnode.Left]
			rvisited, rok := visitedMap[cnode.Right]
			if lok && lvisited && rok && rvisited {
				visitAndSetMap(cnode, visitedMap)
			} else {
				pushBackAdnUpdateVisited(lst, visitedMap, cnode)
				pushBackAdnUpdateVisited(lst, visitedMap, cnode.Right)
				pushBackAdnUpdateVisited(lst, visitedMap, cnode.Left)
			}
		}
	}

	return nil
}
func TransferWidthFirst(root *BinaryTree) error {
	if TreeInvalid(root) {
		return errors.New("input invalid")
	}
	lst := list.New()

	PushBackIfNotNil(lst, root)
	for lst.Len() > 0 {
		cnode := lst.Remove(lst.Front()).(*BinaryTree)
		visit(cnode)
		PushBackIfNotNil(lst, cnode.Left)
		PushBackIfNotNil(lst, cnode.Right)
	}

	return nil
}

///////////////////////////////////util///////////////////////////////
func PushBackIfNotNil(lst *list.List, node *BinaryTree) {
	if node != nil && lst != nil {
		lst.PushBack(node)
	}
}
func pushBackAdnUpdateVisitedIfNotNil(lst *list.List, visitMap map[*BinaryTree]bool, node *BinaryTree) {
	if node == nil {
		return
	}

	pushBackAdnUpdateVisited(lst, visitMap, node)
}
func NodeIsNilOrVisited(node *BinaryTree, visitMap map[*BinaryTree]bool) bool {
	if node == nil {
		return true
	}

	if visited, ok := visitMap[node]; ok && visited {
		return true
	}

	return false
}

func transferMiddleRight(node *BinaryTree, lst *list.List, visitedMap map[*BinaryTree]bool) {
	visitAndSetMap(node, visitedMap)
	if node.Right != nil {
		pushBackAdnUpdateVisited(lst, visitedMap, node.Right)
	}
}
func visit(node *BinaryTree) {
	fmt.Printf("%v ", node.Val)
}
func visitAndSetMap(node *BinaryTree, visitMap map[*BinaryTree]bool) {
	visit(node)
	visitMap[node] = true
}
func pushBackAdnUpdateVisited(lst *list.List, visitedMap map[*BinaryTree]bool, node *BinaryTree) {
	lst.PushBack(node)
	visitedMap[node] = false
}
