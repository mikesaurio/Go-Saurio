package main

import "fmt"

func main(){
	var name string
	fmt.Println("Escribe tu nombre: ")
	fmt.Scan(&name)
	fmt.Println("Hola", name)
}
