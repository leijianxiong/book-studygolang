package main

import "fmt"

/*
1. 解决这个叫做Fizz-Buzz[23]的问题:
编写一个程序，打印从 1 到 100 的数字。当是三的倍数就打印 “Fizz” 代替数字，
当是􏰃5的􏰂数就打印 “Buzz”。当数字同时是三和􏰃的􏰂数 时，打印 “FizzBuzz”。
 */
func main()  {
	for i := 1; i <= 100; i++ {
		fmt.Print(i)
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Printf("[FizzBuzz]")
			break
		case i % 3 == 0:
			fmt.Printf("[Fizz]")
			break
		case i % 5 == 0:
			fmt.Printf("[Buzz]")
			break
		}
		fmt.Print(" ")
		if i % 20 == 0 {
			fmt.Println()
		}
	}
}
