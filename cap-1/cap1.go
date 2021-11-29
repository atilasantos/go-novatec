package cap1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNumbers(c *gin.Context) {
	n1, _ := strconv.Atoi(c.Query("number1"))
	n2, _ := strconv.Atoi(c.Query("number2"))

	c.IndentedJSON(http.StatusOK, n1+n2)
}

func Sum(n1 int, n2 int) int {
	return n1 + n2
}
