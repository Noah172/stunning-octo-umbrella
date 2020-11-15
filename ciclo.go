package main

import (
	"fmt"
)

func main() {

	i:=0

	for {
		fmt.Println(i)
		i++
		if i > 2000 {
			break
		}
	}
}