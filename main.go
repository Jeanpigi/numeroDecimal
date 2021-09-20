package main

import "fmt"


func main() {
	var number int

	fmt.Println("Escribe un número entero: ")
	fmt.Scanf("%d", &number)

	fmt.Println(ContadorDecimal(number))
}

func ContadorDecimal(number int) int {
	decima := number
	contador := 0
	for decima > 0 {
		decima = decima / 10
		contador += 1
	}

	return contador
}