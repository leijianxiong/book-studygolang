package main

import "fmt"

func main() {

	i := 1
	LOOP:
		if i > 10 {
			goto NEXT
		}
		fmt.Println(i)
		i = i + 1
		if i <= 10 {
			goto LOOP
		}
		goto NEXT
	NEXT:
}
