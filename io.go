package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// var nombre string
	// fmt.Println("Escribe tu nombre")
	// fmt.Scanf("%s\n", &nombre)
	// fmt.Printf("Hola %s :D", nombre)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Escribe tu nombre")
	text,err := reader.ReadString('\n')
	
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Printf("Hola "+ text +" :D\n")
	}

}