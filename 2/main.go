package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxBouquetCost(a int) int {
	maxPower := int(0)
	for (1 << maxPower) <= a {
		maxPower++
	}
	maxPower--

	if maxPower < 2 {
		return -1
	}

	for p1 := maxPower; p1 >= 2; p1-- {
		for p2 := p1 - 1; p2 >= 1; p2-- {
			for p3 := p2 - 1; p3 >= 0; p3-- {
				cost := (1 << p1) + (1 << p2) + (1 << p3)
				if cost <= a {
					return cost
				}
			}
		}
	}

	return -1
}

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	amounts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &amounts[i])
	}

	for i := 0; i < n; i++ {
		result := maxBouquetCost(amounts[i])
		fmt.Println(result)
	}
}
