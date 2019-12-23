package main

import (
	"fmt"
	"strings"
)

func main()  {
	f := func(i int) int {
		return i * 2
	}
	a := []int{1, 2, 3, 4}

	newa := m1(f, a)
	fmt.Println("a", a)
	fmt.Println("newa", newa)


	//
	f2 := func(i interface{}) interface{} {
		v, ok := i.(string)
		if !ok {
			panic(fmt.Sprintf("i is not string: %v", i))
		}
		return strings.ToUpper(v)
	}
	b := []interface{}{"a", "b", "c", "c"}

	newb := m2(f2, b)
	fmt.Println("b", b)
	fmt.Println("newb", newb)
}

//支持int
func m1(f func(i int) int, a []int) (newa []int) {
	newa = make([]int, len(a))

	for k, v := range a {
		newa[k] = f(v)
	}

	return newa
}

//支持string
func m2(f func(i interface{}) interface{}, a []interface{}) (newa []interface{}) {
	newa = make([]interface{}, len(a))

	for k, v := range a {
		newa[k] = f(v)
	}

	return newa
}
