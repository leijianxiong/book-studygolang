package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var f = float64(20) / float64(3)
	//其他类型转字符串
	var fs = fmt.Sprintf("%v", f)
	fmt.Println("f[string]=>", fs)

	// 字符串转其他类型 参考strconv
	fsTof, _ := strconv.ParseFloat(fs, 64)
	fmt.Println("f[float]=>", fsTof)
}
