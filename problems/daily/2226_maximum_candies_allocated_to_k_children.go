package daily

func maximumCandies(candies []int, k int64) int {
	left, right := 1, 10_000_000
	result := 0

	for left <= right {
		mid := left + (right-left)/2

		var avgCandies int
		for _, candy := range candies {
			avgCandies += candy / mid
		}

		if int64(avgCandies) < k {
			right = mid - 1
		}

		if k <= int64(avgCandies) {
			result = mid
			left = mid + 1
		}
	}

	return result
}
