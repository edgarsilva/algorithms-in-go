package main

import (
	"fmt"
	"math"
)

func main() {
	input := "abc"
	result := cartesianProduct(input, 3)

	fmt.Printf("Cartesian Product ->\n")
	for _, elem := range result {
		fmt.Printf("%v\n", elem)
	}
	fmt.Printf("Product length %v\n", len(result))
}

func cartesianProduct(chars string, n int) []string {
	if n <= 0 {
		return []string{""}
	}

	smallerProduct := cartesianProduct(chars, n-1)

	base := float64(len(chars))
	exp := float64(n)
	resultLen := int64(math.Pow(base, exp))

	result := make([]string, 0, resultLen)
	for _, prefix := range smallerProduct {
		for _, char := range chars {
			result = append(result, prefix+string(char))
		}
	}
	return result
}
