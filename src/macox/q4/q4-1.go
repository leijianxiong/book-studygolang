package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
编写计算一个类型是 float64 的 slice 的平均值的代码。在稍候的练习 Q5 中 将会改写为函数
 */
func main()  {
	rand.Seed(time.Now().UnixNano())
	//生成1到10随机数
	min := 5
	max := 10
	randNum := rand.Intn(max - min) + min

	sf := make([]float64, randNum)
	var sum float64

	for i := 0; i <= randNum - 1; i++ {
		value := rand.Float64() * float64(randNum)
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
		sf[i] = value
		sum += value
	}

	fmt.Println(len(sf), sf)

	//计算平均值
	var svg = sum / float64(randNum)

	fmt.Println("sum", sum)
	fmt.Println("svg", svg)

	sum, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", sum), 64)
	svg, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", svg), 64)
	fmt.Println("sum .2f", sum)
	fmt.Println("svg .2f", svg)

}
