package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
创建一个固定大小保存整数的栈。它无须超出限制的增长。
定义push函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出(LIFO) 的。

1-2
 */

type Stack struct {
	data []int
	p int
}

func NewStack() *Stack {
	return &Stack{
		//data: make([]int, 1),
		data: *new([]int),
		p: 0,
	}
}

func (s *Stack) Push(e int) {
	s.data = append(s.data, e)
	s.p += 1
}

func (s *Stack) Pop() (e int)  {
	if s.p == 0 {
		panic("data empty")
	}
	s.p -= 1

	e = s.data[s.p]
	s.data[s.p] = 0
	return
}

func (s *Stack) String() string {
	if s.p < 1 {
		return ""
	}

	str := fmt.Sprintf("len:%d", s.p)
	for k, v := range s.data[0: s.p] {
		str += fmt.Sprintf(" [%d:%d]", k, v)
	}

	return str
}

func main()  {
	s := NewStack()
	fmt.Println("stack", s.String())

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 3; i++ {
		e := int(rand.Float64() * 100)
		s.Push(e)
	}

	fmt.Println("stack", s.String())
	fmt.Println("stack pop", s.Pop())
	fmt.Println("stack", s.String())
	fmt.Printf("My stack %v \n", s)
}