package main

import (
	"fmt"
	"math"
)

func main() {
	//declaração de variável
	var angulo float64

	//entrada de dado
	fmt.Println("Digite o ângulo desejado: ")
	fmt.Scan(&angulo)

	//cálculo e impressão do cosseno
	fmt.Println("Cosseno de", angulo, "é:", math.Cos(angulo))

	//cálculo e impressão do seno
	fmt.Println("Seno de", angulo, "é:", math.Sin(angulo))
}
