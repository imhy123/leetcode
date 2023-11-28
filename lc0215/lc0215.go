package lc0215

var Target = findKthLargest

func adjustUp(tree []int, index int) {
	parentIdx := (index - 1) / 2

	if tree[index] > tree[parentIdx] {
		t := tree[parentIdx]
		tree[parentIdx] = tree[index]
		tree[index] = t

		if parentIdx > 0 {
			adjustUp(tree, parentIdx)
		}
	}
}

func adjustDown(tree []int, size int, index int) {
	leftIdx := 2*index + 1
	if leftIdx >= size {
		return
	}

	targetIdx := index
	if tree[leftIdx] > tree[index] {
		targetIdx = leftIdx
	}

	rightIdx := 2*index + 2
	if rightIdx < size && tree[rightIdx] > tree[targetIdx] {
		targetIdx = rightIdx
	}

	if targetIdx == index {
		return
	}

	t := tree[index]
	tree[index] = tree[targetIdx]
	tree[targetIdx] = t

	adjustDown(tree, size, targetIdx)
}

func insertToTree(tree []int, size int, val int) int {
	tree[size] = val

	if size == 0 {
		return val
	}

	adjustUp(tree, size)

	return tree[size]
}

func getTop(tree []int, size int) int {
	ret := tree[0]
	tree[0] = tree[size-1]

	adjustDown(tree, size-1, 0)

	return ret
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	var tree []int = make([]int, length)

	for i := 0; i < length; i++ {
		insertToTree(tree, i, nums[i])
	}

	ret := 0
	for i := 0; i < k; i++ {
		ret = getTop(tree, length-i)
	}

	return ret
}
