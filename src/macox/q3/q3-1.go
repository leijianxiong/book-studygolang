package main

import "fmt"

/*
建立一个Go程序打印下面的内容(到100个字符):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
 */
func main()  {
	for lineNum := 1; lineNum <= 100; lineNum++ {
		for charNum := 1; charNum <= lineNum; charNum++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}
