package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
编写一个函数用于计算一个float64类型的slice的平均值。
 */
func main()  {
	//生成一个[]float64

	rand.Seed(time.Now().UnixNano())
	min := 2
	max := 5
	randNum := rand.Intn(max - min) + min

	generateRandomFloat64 := func() float64 {
		return rand.Float64() * float64(100)
	}

	sf := make([]float64, randNum)

	for k, _ := range sf {
		sf[k] = generateRandomFloat64()
	}

	fmt.Println("sf", randNum, sf)
	fmt.Println("svg:", demo1(sf))
}

func demo1(sf []float64) (svg float64)  {
	sum := float64(0)
	for _, v := range sf {
		sum += v
	}

	fmt.Println("sum", sum)
	svg = sum / float64(len(sf))
	return
}
