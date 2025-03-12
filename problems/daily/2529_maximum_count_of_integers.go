package daily

func maximumCount(nums []int) int {
	var (
		negCount int
		posCount int
	)

	for _, n := range nums {
		if n == 0 {
			continue
		}

		if n < 0 {
			negCount++
		}

		if n > 0 {
			posCount++
		}
	}

	switch {
	case negCount > posCount:
		return negCount
	case posCount > negCount:
		return posCount
	default:
		return posCount
	}
}
