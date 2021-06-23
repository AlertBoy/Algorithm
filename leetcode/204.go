package main

import "fmt"

func countPrimes(n int) int {
	var temp []bool = make([]bool, n)
	result := 0

	for i := 2; i < n; i++ {
		if !temp[i] {
			result++
			for j := i * 2; j < n; j += i {
				temp[j] = true
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countPrimes(12))
}
