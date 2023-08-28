func twoSum(nums []int, target int) []int {
	past := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		d := target - n
		j, has := past[d]
		if has {
			return []int{i, j}
		}
		past[n] = i
	}
	return []int{0, 0}
}
