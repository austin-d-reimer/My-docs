package main

import "fmt"

func main() {
	maxnum := 100000
	fmt.Println("Hello")
	count := 1
	for i := 1; i <= maxnum; i++ {
		if i%2 != 0 {
			if isPrime(i) {
				count++
			}
		}
	}
	fmt.Println("There are", count, "many prime numbers. in the first", maxnum, "numbers")
}

func isPrime(num int) bool {

	i := 3
	for i < num {
		if num%i == 0 {
			return false
		}
		i += 2
	}
	return true
}
