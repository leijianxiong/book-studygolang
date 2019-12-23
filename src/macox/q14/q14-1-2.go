package main

import "fmt"

/*
(1) 函数返回一个函数
1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名
称叫做 plusTwo。然后可以像下面这样使用: p := plusTwo()
      fmt.Printf("%v\n", p(2))
应该打印 4。参阅第 31 页的 “回调” 小节了解更多相关信息。
2. 使1中的函数更加通用化，创建一个plusX(x)函数，返回一个函数用于对整 数加上 x。
*/
func main() {
	//1
	p := plusTwo()

	fmt.Printf("%v\n", p(2))

	//2
	q := plusX(100)
	fmt.Printf("%v\n", q(10))
}

func plusTwo() (f func(i int) int) {
	return func(i int) int {
		return i + 2
	}
}

func plusX(x int) (f func(i int) int) {
	return func(i int) int {
		return x + i
	}
}
