package main

import "fmt"

func eliminarSumaCero(arr []int) []int {
	stack := []int{}

	for _, num := range arr {
		stack = append(stack, num)
		// Buscamos si hay subsecuencias cuya suma sea 0
		sum := 0
		for i := len(stack) - 1; i >= 0; i-- {
			sum += stack[i]
			if sum == 0 {
				// Si la suma es 0, eliminamos esa subsecuencia
				stack = stack[:i]
				break
			}
		}
	}

	return stack
}

func main() {
	arr := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	fmt.Println(eliminarSumaCero(arr)) // [5, 8]
}
