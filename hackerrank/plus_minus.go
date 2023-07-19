package main

import (
	"fmt"
)

func main() {
	var n, pos, neg, zeros, t int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t > 0 {
			pos++
		} else if t < 0 {
			neg++
		} else {
			zeros++
		}
	}
	pf := float32(pos) / float32(n)
	nf := float32(neg) / float32(n)
	zf := float32(zeros) / float32(n)

	fmt.Printf("%.6f\n", pf)
	fmt.Printf("%.6f\n", nf)
	fmt.Printf("%.6f\n", zf)
}
