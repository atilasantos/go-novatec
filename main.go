package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	cap1 "github.com/atilasantos/go-novatec/cap-1"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	//Soma de duas entradas
	fmt.Println("\n\nSomando atributos")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(cap1.Sum(n1, n2))

	//Cap1 - Exercicio1.1 exibindo todos os parâmetros
	fmt.Println("\n\nExibindo todos argumentos concatenados com - ")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	start1 := time.Now()
	fmt.Println(strings.Join(os.Args, "-"))
	end1 := time.Now()

	//Cap1 - Exercicio1.2
	fmt.Println("\n\nExibindo índice e valor de cada argumento")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	start2 := time.Now()
	for index, arg := range os.Args {
		fmt.Printf("Índice: %v\nargumento: %v\n", index, arg)
	}
	end2 := time.Now()

	//Cap1 - Exercicio1.3
	fmt.Println("\nExibindo diferença de tempo de execução:")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("Tempo utilizando join: %v\nTempo utilizando for: %v", end1.Sub(start1), end2.Sub(start2))

	// router := gin.Default()
	// router.GET("/sum", cap1.GetNumbers)
	// router.Run("localhost:8000")
}
