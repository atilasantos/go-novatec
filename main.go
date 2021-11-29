package main

import (
	"fmt"
	"os"
	"strconv"

	cap1 "github.com/atilasantos/go-novatec/cap-1"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	fmt.Println(cap1.Sum(n1, n2))
}
