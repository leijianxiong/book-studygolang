package main

import (
	"fmt"
)

/*
1. 编写一个函数，找到intslice([]int)中的最大值。
2. 编写一个函数，找到intslice([]int)中的最小值
 */
func main()  {
	ints := []int{489 , 9308, 18 , 48,7 , 1209 , 1, -423, -4999, -42139}
	//ints := []int{489 , 9308, 18 , 48,7 , 1209 , 1}

	fmt.Println("ints", ints)
	fmt.Println("max", max(ints))
	fmt.Println("min", min(ints))
}

func max(a []int) int {
	r := a[0]

	for i := 1; i <= len(a) - 1; i++ {
		if a[i] > r {
			r = a[i]
		}
	}

	return r
}

func min(a []int) int {
	r := a[0]

	for i := 1; i <= len(a) - 1; i++ {
		if a[i] < r {
			r = a[i]
		}
	}

	return r
}
