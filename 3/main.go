package main

import (
	"fmt"
	"sort"
)

func goodDays(n, m int, a []int) int {
	goodDays := 0

	for i := 2; i < n; i++ {
		if a[i] >= a[0] && a[i] <= a[1] {
			goodDays++
		}
	}

	if goodDays >= m {
		return 0
	}

	badDays := make([]int, 0)
	for i := 2; i < n; i++ {
		if !(a[i] >= a[0] && a[i] <= a[1]) {
			badDays = append(badDays, a[i])
		}
	}

	sort.Ints(badDays)

	adjustments := 0
	for i := 0; i < len(badDays) && goodDays < m; i++ {

		if badDays[i] < a[0] {
			adjustments += a[0] - badDays[i]
		} else if badDays[i] > a[1] {
			adjustments += badDays[i] - a[1]
		}
		goodDays++
	}

	return adjustments
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	result := goodDays(n, m, a)
	fmt.Println(result)
}
