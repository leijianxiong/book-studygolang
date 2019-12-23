package main

import "fmt"

/*
编写函数接受整数类型变参，并且每行打印一个数字。
 */

func main() {
	demo1(1, 3, 4, 5)
}

func demo1(arg ...int)  {
	for k, v := range arg {
		fmt.Println(k, v)
	}
}