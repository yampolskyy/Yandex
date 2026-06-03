package main

import (
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)
	ans := append(left, right...)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ans[k] = left[i]
			k++
			i++
		} else {
			k++
			j++
		}
	}
	return ans
}
