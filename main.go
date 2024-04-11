package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	var n int64
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	fmt.Printf("Factorial of %d is: %v\n", n, factorial(n))
}
