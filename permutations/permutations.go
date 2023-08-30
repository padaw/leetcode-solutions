// I don't know go
func remove(arr []int, idx int) []int {
	size := len(arr)
	next := make([]int, size-1)
	done := 0
	for i := 0; i < size && done < size-1; i++ {
		if i == idx {
			continue
		}
		next[done] = arr[i]
		done++
	}
	return next
}

func permute(nums []int) [][]int {
	size := len(nums)

	if size < 2 {
		return [][]int{nums}
	}
	if size == 2 {
		return [][]int{nums, {nums[1], nums[0]}}
	}

	res := [][]int{}

	for i := 0; i < size; i++ {
		n := nums[i]
		next := remove(nums, i)
		list := permute(next)
		for i := 0; i < len(list); i++ {
			out := append([]int{n}, list[i]...)
			res = append(res, out)
		}
	}

	return res
}
