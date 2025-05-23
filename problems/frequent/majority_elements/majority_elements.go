package majority_elements

func majorityElement(nums []int) int {
	countDict := make(map[int]int, 0)
	majority := nums[0]

	for _, majority = range nums {
		if _, ok := countDict[majority]; ok {
			countDict[majority] += 1
		} else {
			countDict[majority] = 1
		}
	}
	for k, v := range countDict {
		if v > len(nums)/2 {
			majority = k
		}
	}
	return majority
}
