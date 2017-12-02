package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x,y := 0,0
	return func() int {
		if x == 0 && y == 0 {
			y++
			return 0
		}

		x = x + y
		y = x - y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
