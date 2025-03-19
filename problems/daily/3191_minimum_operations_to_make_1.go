package daily

func minOperations(nums []int) int {
	l := len(nums)

	if l < 3 {
		return -1
	}

	var count int

	for i := 0; i < l-2; i++ {
		if nums[i] == 0 {
			nums[i] = flip(nums[i])
			nums[i+1] = flip(nums[i+1])
			nums[i+2] = flip(nums[i+2])
			count++
		}
	}

	if nums[l-2] == 1 && nums[l-1] == 1 {
		return count
	}

	return -1
}

func flip(e int) int {
	if e == 0 {
		return 1
	}
	return 0
}
