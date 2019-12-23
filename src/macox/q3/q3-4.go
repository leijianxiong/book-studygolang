package main

import "fmt"

/*
编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。
提示:不 幸的是你需要知道一些关于转换的内容，参阅 “转换” 第 59 页的内容。
 */
func main()  {
	//demo1()
	demo2()
}

func demo1()  {
	str := "abcdef"
	strRunes := []rune(str)
	runeLen := len(strRunes)

	//这种方法会带空格
	//revertRunes := make([]rune, runeLen)
	revertRunes := *new([]rune)

	for i := 1; i <= runeLen; i++ {
		lastIndex := runeLen - i
		revertRunes = append(revertRunes, strRunes[lastIndex])
	}

	fmt.Printf("run-len:%d, %s '%s'\n", runeLen, str, string(revertRunes))
}

func demo2()  {
	//str := "abcdef"
	str := "发多了发哈圣诞恋歌"
	strRunes := []rune(str)
	runeLen := len(strRunes)

	revertRunes := make([]rune, runeLen)

	for i, maxIndex := 0, runeLen - 1; i <= maxIndex; i++ {
		lastIndex := maxIndex - i
		revertRunes[lastIndex] = strRunes[i]
	}

	fmt.Printf("run-len:%d, %s '%s'\n", runeLen, str, string(revertRunes))
}
