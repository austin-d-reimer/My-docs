package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var b bool = false
	// testing how variables with the same name works
	x := 8
	fmt.Println(x)
	fmt.Println(add(18, 21))
	fmt.Println(x)

	fmt.Println(swap("World!", "Hello"))

	loop(5)

	fmt.Println(iforelse(b))
}

func add(x int, y int) int {
	return x + y
}

// returning more than one variable at a time
func swap(a string, b string) (string, string) {
	return b, a
}

// running a for loop
func loop(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Println(i)
	}
}

func iforelse(b bool) string {
	if b {
		return "b is true"
	} else {
		return "b is false"
	}

}
