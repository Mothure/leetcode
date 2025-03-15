package daily

func minCapability(nums []int, k int) int {
	lo, hi := minInArray(nums), maxInArray(nums)

	for lo < hi {
		mid := (lo + hi) / 2
		if canRob(nums, k, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canRob(nums []int, k, capability int) bool {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= capability {
			count++
			i++
		}
		i++
	}
	return count >= k
}

func minInArray(arr []int) int {
	minVal := arr[0]
	for _, num := range arr {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func maxInArray(arr []int) int {
	maxVal := arr[0]
	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}
