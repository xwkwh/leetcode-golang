package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))

}
func hammingDistance(x int, y int) int {
	var res int
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			res ++
		}
		x /= 2
		y /= 2
	}
	return res
}
