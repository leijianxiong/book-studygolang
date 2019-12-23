package main

import (
	"fmt"
	"unicode/utf8"
)

/*
建立一个程序统计字符串里的字符数量:
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。提示: 看看 unicode/utf8 包
 */
func main() {
	//s := "asSASA ddd dsjkdsjs dk"
	s := "Hello, 世界"

	count := utf8.RuneCountInString(s)
	fmt.Println("count: ", count)

	for k, v := range s {
		_ = k
		fmt.Println(string(v), len(string(v)))
	}
}
