package main

import (
	"fmt"
	"math"
)

func main() {
	//var arr [][]int32

	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var sum_diag int32
	var sum_diag_inv int32

	for i := 0; i < len(arr); i++ {
		sum_diag += arr[i][i]
		sum_diag_inv += arr[i][len(arr)-1-i]
	}

	res := sum_diag - sum_diag_inv
	abs := math.Abs(float64(res))

	fmt.Println(int32(abs))
}
