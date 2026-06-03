package main

import (
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)
	ans := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}

	for i < len(left) {
		ans = append(ans, left[i])
		i++
	}

	for j < len(right) {
		ans = append(ans, right[j])
		j++
	}

	return ans
}
