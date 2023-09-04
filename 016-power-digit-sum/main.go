package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// number as string
	num := math.Pow(2, 1000)
	numChar := fmt.Sprintf("%.0f", num)
	digitsChar := strings.Split(numChar, "")
	digitsInt := make([]int, len(digitsChar))
	for i := 0; i < len(digitsChar); i++ {
		digitsInt[i], _ = strconv.Atoi(digitsChar[i])
	}
	total := 0
	for i := 0; i < len(digitsInt); i++ {
		total += digitsInt[i]
	}
	fmt.Println(total)
}
