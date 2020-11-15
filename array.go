package main

import (
	"fmt"
)

func main() {

	arry := [4]int {4, 15, 2, 12}

	for i := 0; i < len(arry); i++ {
		fmt.Println(arry[i])
	}

	// matrices

	var mat [3][3]int
	fmt.Println("============ esto es una matriz============")
	for x := 0; x < len(mat); x++ {
		fmt.Println(mat[x])
	}

}