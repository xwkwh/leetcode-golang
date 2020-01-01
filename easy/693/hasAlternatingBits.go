package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5) == true)
	fmt.Println(hasAlternatingBits(7) == false)
	fmt.Println(hasAlternatingBits(11) == false)
	fmt.Println(hasAlternatingBits(10) == true)
	fmt.Println(hasAlternatingBits(4) == false)
	//fmt.Println(0 ^ 0 )
}


func hasAlternatingBits(n int) bool {
	pre := n & 1
	cur := 0
	for n > 0 {
		if cur == 0 {
			cur ++
			n >>= 1
			continue
		}
		bit := n & 1
		if bit + pre != 1 {
			return false
		}
		pre = bit
		n >>= 1
	}
	return true
}