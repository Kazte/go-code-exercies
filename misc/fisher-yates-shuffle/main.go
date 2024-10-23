package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array)
	shuffle(array)
	fmt.Println(array)
}

func shuffle(arr []int) []int {
	fmt.Println("Shuffling")
	for i := len(arr) - 1; i > 1; i-- {
		az := rand.Intn(i)

		tmp := arr[az]
		arr[az] = arr[i]
		arr[i] = tmp
	}

	return arr
}
