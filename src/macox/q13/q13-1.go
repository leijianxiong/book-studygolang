package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
编写一个针对int类型的slice冒泡排序的函数。这里[24]:
它在一个列表上重复步骤来排序，比较每个相􏰄的元素，并且顺序错 误的时候，交换它们。
一遍一遍扫描列表，直到没有交换为止，这意 味着列表排序完成。算法得名于更小的元素就像 “泡泡” 一样冒到列表 的􏰅端。
 */
func main()  {
	rand.Seed(time.Now().UnixNano())

	num := int(math.Ceil(rand.Float64() * 10))
	a := make([]int, num)

	for i, _ := range a {
		a[i] = int(rand.Float64() * 300)
	}

	fmt.Printf("a len:%d =>%v\n", num, a)

	bubbleSort(a)
}

//冒泡排序 时间复杂度 O((n-1)^2)..
func bubbleSort(a []int)  {
	aCount := len(a) - 1

	for l := 1; l <= aCount; l++ {
		for i := 0; i <= aCount - 1; i++ {
			x := a[i]
			y := a[i+1]
			if x > y {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}

		fmt.Println("a", l, a)
	}

	fmt.Println("r a", a)
}