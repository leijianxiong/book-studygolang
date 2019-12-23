package main

import "fmt"

func main()  {
	var a []int
	b := *new([]int)
	fmt.Println(a, b)
}
