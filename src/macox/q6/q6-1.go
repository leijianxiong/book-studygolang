package main

import (
	"fmt"
	"sort"
)

/*
编写函数，返回其(两个)参数正确的(自然)数字顺序:
 */
func main()  {
	demo1(7, 2, 3)
	demo1(2, 7, 4)
}

func demo1(arg ...int)  {
	fmt.Println("sort Before", arg)
	sort.Ints(arg)
	fmt.Println("sort After", arg)
}
