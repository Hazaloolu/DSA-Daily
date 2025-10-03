package main

import "fmt"

func FindSumPair(n int) [][2]int {

	pairs := make([][2]int, 0, n+1)

	for a := 0; a <= n; a++ {

		b := n - a

		pairs = append(pairs, [2]int{a, b})

	}
	return pairs

}

func main() {
	targetN := 5

	result := FindSumPair(targetN)

	fmt.Printf("The pair of numbers that sum up to %d are \n", targetN)

	for _, pair := range result {

		fmt.Printf("(%d, %d)\n", pair[0], pair[1])

	}
}
