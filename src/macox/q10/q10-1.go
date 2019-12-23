package main

import "fmt"

func main()  {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d fib(%d)=%d\n", i, i, fib(i))
	}
}

func fib(i int) int {
	if i < 1 {
		panic(fmt.Sprintf("invalid i: %d", i))
	}

	if i == 1 || i == 2 {
		return 1
	}

	return fib(i - 1) + fib(i - 2)
}
