package main

import (
	"bufio"
	"fmt"
	"os"
)

const CONST_VALUE = 998244353

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &array[i])
	}

	results := make([]int, k)
	for p := 1; p <= k; p++ {

		sums := make([]int, 0, n*(n-1)/2)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				sums = append(sums, array[i]+array[j])
			}
		}

		totalSum := 0
		for _, v := range sums {
			totalSum = (totalSum + PowerMod(v, p, CONST_VALUE)) % CONST_VALUE
		}

		results[p-1] = totalSum
	}

	for _, result := range results {
		fmt.Fprintln(out, result)
	}
}

func PowerMod(base, exp, mod int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}
