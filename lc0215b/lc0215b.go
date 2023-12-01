package lc0215b

var Target = findKthLargest

type AvlNode struct {
	Val        int
	LeftChild  *AvlNode
	RightChild *AvlNode

	ValCount   int
	LeftCount  int
	RightCount int
}

func insertToTree(root *AvlNode, val int) {
	if val == root.Val {
		root.ValCount = root.ValCount + 1
		return
	}

	if val < root.Val {
		if root.LeftChild == nil {
			newNode := new(AvlNode)
			newNode.Val = val
			newNode.ValCount = 1

			root.LeftChild = newNode
			root.LeftCount = 1
			return
		} else {
			insertToTree(root.LeftChild, val)
			root.LeftCount = root.LeftCount + 1
		}
	} else {
		if root.RightChild == nil {
			newNode := new(AvlNode)
			newNode.Val = val
			newNode.ValCount = 1

			root.RightChild = newNode
			root.RightCount = 1
			return
		} else {
			insertToTree(root.RightChild, val)
			root.RightCount = root.RightCount + 1
		}
	}
}

func findKthInTree(root *AvlNode, k int) int {
	curNode := root
	curK := k

	for {
		if curNode.RightCount >= curK {
			curNode = curNode.RightChild
			continue
		}

		if curNode.RightCount+curNode.ValCount >= curK {
			return curNode.Val
		}

		curK = curK - curNode.RightCount - curNode.ValCount
		curNode = curNode.LeftChild
	}
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)

	if length == 1 {
		return nums[0]
	}

	var root *AvlNode = new(AvlNode)
	root.Val = nums[0]
	root.ValCount = 1

	for i := 1; i < length; i++ {
		insertToTree(root, nums[i])
	}

	return findKthInTree(root, k)
}
