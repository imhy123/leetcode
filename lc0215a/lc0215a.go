package lc0215a

var Target = findKthLargest

func findKthLargest(nums []int, k int) int {
	length := len(nums)

	if length == 1 {
		return nums[0]
	}

	idx := length / 2
	idxVal := nums[idx]
	equalsCount := 0

	var left []int
	var right []int

	for i := 0; i < length; i++ {
		if i == idx {
			continue
		}

		if nums[i] < idxVal {
			left = append(left, nums[i])
		} else if nums[i] == idxVal {
			equalsCount++
		} else {
			right = append(right, nums[i])
		}
	}

	lenRight := len(right)

	if lenRight >= k {
		return findKthLargest(right, k)
	} else if lenRight+equalsCount >= k-1 {
		return idxVal
	} else {
		return findKthLargest(left, k-lenRight-equalsCount-1)
	}
}
