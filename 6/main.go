package main

import (
	"bufio"
	"fmt"
	"os"
)

func Evklid(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	if n < 3 {
		fmt.Fprintln(out, 0)
		return
	}

	coords := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &coords[i][0], &coords[i][1])
	}

	L := 1

	for i := 0; i < n; i++ {
		slopesCount := make(map[[2]int]int)
		x1, y1 := coords[i][0], coords[i][1]

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			x2, y2 := coords[j][0], coords[j][1]
			dx, dy := x2-x1, y2-y1

			if dx == 0 {

				slopesCount[[2]int{0, 1}]++
			} else {
				sign := 1
				if dx*dy < 0 {
					sign = -1
				}

				dx, dy = abs(dx), abs(dy)
				g := Evklid(dx, dy)
				dx /= g
				dy /= g

				slopesCount[[2]int{sign * dy, dx}]++
			}
		}

		localMax := 0
		for _, count := range slopesCount {
			if count > localMax {
				localMax = count
			}
		}

		L = max(L, localMax+1)
	}

	if float64(L) <= float64(2*n)/3 {
		fmt.Fprintln(out, n/3)
	} else {
		fmt.Fprintln(out, n-L)
	}
}
