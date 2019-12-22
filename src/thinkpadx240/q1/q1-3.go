package main

import "fmt"

func main() {
	arr1 := [...]int{1,2,3,4,5,6,7,8,9,10}
	for k, v := range arr1 {
		fmt.Println("ev arr element: ", k, v)
	}
}
