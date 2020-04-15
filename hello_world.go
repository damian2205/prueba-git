package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Ingresa tu nombre")
	fmt.Scan(&nombre)
	fmt.Print("HOLA  " + nombre + "  :)")
}
