package main

import "fmt"

func main() {
	moves := "LLRR"
	fmt.Println(judgeCircle(moves))
}
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, c := range []rune(moves) {
		switch string(c) {
		case "L":
			x++
		case "R":
			x--
		case "U":
			y ++
		case "D":
			y --
		}
	}
	return x == 0 && y == 0
}
