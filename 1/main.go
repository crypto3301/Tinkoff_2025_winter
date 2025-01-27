package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkStr(line string) bool {
	r_ind := -1
	m_ind := -1

	for i, ch := range line {
		if ch == 'R' {
			r_ind = i
		} else if ch == 'M' {
			m_ind = i
		}
	}

	return r_ind < m_ind
}

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	if checkStr(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
