package main

import (
	"fmt"
)

func main() {
	var n, x, y, z int
	fmt.Scan(&n, &x, &y, &z)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	minOps := 0
	minX, minY, minZ := -1, -1, -1
	opsX, opsY, opsZ := 0, 0, 0

	for i := 0; i < n; i++ {

		remX := (x - a[i]%x) % x
		remY := (y - a[i]%y) % y
		remZ := (z - a[i]%z) % z

		if minX == -1 || remX < (x-a[minX]%x)%x {
			minX = i
			opsX = remX
		}
		if minY == -1 || remY < (y-a[minY]%y)%y {
			minY = i
			opsY = remY
		}
		if minZ == -1 || remZ < (z-a[minZ]%z)%z {
			minZ = i
			opsZ = remZ
		}
	}

	minOps = opsX + opsY + opsZ

	if minX == minY && minY == minZ {
		minOps = max(opsX, max(opsY, opsZ))
	} else if minX == minY {
		minOps = max(opsX, opsY) + opsZ
	} else if minX == minZ {
		minOps = max(opsX, opsZ) + opsY
	} else if minY == minZ {
		minOps = max(opsY, opsZ) + opsX
	}

	fmt.Println(minOps)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
