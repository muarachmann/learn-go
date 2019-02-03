package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		println(a[i])
	}

	randomGenerateNumber := func(num int) int {
		var number int = 0
		number = rand.Intn(num)
		return number
	}

	// two dimensional arrays
	var twoD [5][5]int
	fmt.Println(twoD)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = randomGenerateNumber(20)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(twoD[i][j])
		}
	}

}
