package lc

func MaxSubArray(nums []int) int {
	maxSoFar := -1 << 31
	maxEndingHere := 0

	for i := 0; i < len(nums); i++ {
		maxEndingHere += nums[i]

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return maxSoFar
}
