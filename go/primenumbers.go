package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	run_timed()
}

func run(maxnum int) int {
	rawbits := make([]bool, int((maxnum+1)/2))
	count := 0
	for f := 3; f < int(math.Sqrt(float64(maxnum))+1); f += 2 {
		for num := f; num < maxnum; num += 2 {
			if num%2 != 0 && rawbits[int(num/2)] == false {
				f = num
				break
			}
		}

		for number := f * 3; number < maxnum; number += f * 2 {
			rawbits[int(number/2)] = true
		}
	}

	for i := 0; i < len(rawbits); i++ {
		if !rawbits[i] {
			count++
		}
	}
	return count
}

func run_timed() {
	starttime := time.Now()
	count := 0
	result := 0
	for {
		result = run(1_000_000)
		count++
		if result != 78_498 {
			fmt.Println(result, " Is not correct")
		}
		if time.Now().Second()-starttime.Second() > 5 {
			break
		}
	}
	fmt.Println("Ran ", count, " times in 5 seconds with the result of ", result)
}
