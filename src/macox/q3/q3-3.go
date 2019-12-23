package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
扩展/修改上一个问题的程序，替换位置4开始的三个字符为“abc”。

拷贝0,4 abc, (4+3 - -1)
 */
func main()  {
	//str := "Hello, 世界"
	str := "你好佛阿仕顿法拉盛大法师"
	//str := "asSASA ddd dsjkdsjs dk"
	pos := 4
	replaceStr := "abc"

	newstr := ReplaceN(str, pos, replaceStr)

	fmt.Println("1 old str:", str)
	fmt.Println("1 new str:", newstr)

	fmt.Println("2 old str:", str)
	fmt.Println("2 new str:", ReplaceN2(str, pos, replaceStr))

	/////////
	s := str
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After : %s\n", string(r))
}

// pos为字节数
func ReplaceN(str string, pos int, replaceStr string) string {
	var oldstrBuffer = bytes.NewBufferString(str)
	var newstrBuffer bytes.Buffer

	io.CopyN(&newstrBuffer, oldstrBuffer, int64(pos))
	io.CopyN(&newstrBuffer, bytes.NewBufferString(replaceStr), int64(len(replaceStr)))


	for i := 1; i <= len([]rune(replaceStr)); i++ {
		oldstrBuffer.ReadRune()
	}

	io.CopyN(&newstrBuffer, oldstrBuffer, int64(oldstrBuffer.Len()))


	return newstrBuffer.String()
}

//pos为字符数
func ReplaceN2(str string, pos int, replaceStr string) string  {
	strs := []rune(str)
	newstrRunes := *new([]rune)
	replaceStrRunes := []rune(replaceStr)

	newstrRunes = append(newstrRunes, strs[0:pos]...)
	newstrRunes = append(newstrRunes, replaceStrRunes...)

	pos += len(replaceStrRunes)
	lastRunes := strs[pos:]
	newstrRunes = append(newstrRunes, lastRunes...)

	return string(newstrRunes)
}
