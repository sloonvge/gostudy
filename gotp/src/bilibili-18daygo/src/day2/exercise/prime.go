package main

import (
	"math"
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
	var prime bool
	prime = true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			prime =  false
			break
		}
	}
	return prime
}

func isNarcissistic(n int) bool {
	nString := fmt.Sprintf("%s", n)
	var sum float64
	for _, s := range nString {
		sStr := string(s)
		sInt, _ := strconv.ParseInt(sStr, 10, 32)
		sum += math.Pow(float64(sInt), 3)
	}
	if sum == float64(n) {
		return true
	}
	return false
}


func factorialSum(n int) int {
	var result int
	var factorial = 1
	for i := 1; i < n + 1; i++ {
		factorial = factorial * i
		result += factorial
	}
	return result
}

func main() {
	for i := 100; i < 200; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	for i := 100; i < 1000; i++ {
		if isNarcissistic(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	fmt.Printf("1! + 2! + 3! = %d\n", factorialSum(3))
}