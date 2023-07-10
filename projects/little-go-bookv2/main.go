package main;

import (
	"sort"
)

func main() {
	scores := []int{1,2,3,4,5,6,7,8,9,10}
	ind := binarySearch(scores, 6)
	println(scores[ind])
}

func binarySearch(arr []int, target int) (int) { // bayza a7a
	sort.Ints(arr)
	s := 0
	e := len(arr) - 1
	m := 0
	for s <= e {
		m = (s + e) / 2
		if arr[m] == target {
			return s
		} else if arr[m] < target  {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}