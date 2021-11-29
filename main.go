package main

import (
	cap1 "github.com/atilasantos/go-novatec/cap-1"
	"github.com/gin-gonic/gin"
)

func main() {
	// n1, _ := strconv.Atoi(os.Args[1])
	// n2, _ := strconv.Atoi(os.Args[2])

	// fmt.Println(cap1.Sum(n1, n2))

	router := gin.Default()
	router.GET("/sum", cap1.GetNumbers)
	router.Run("localhost:8000")
}
