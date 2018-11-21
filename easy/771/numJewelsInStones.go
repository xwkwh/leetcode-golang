package main

import "fmt"

func main() {
	J := "aA"
	S := "aAAsdfs"
	fmt.Println(numJewelsInStones(J, S))
	for _, char := range []rune(J) {
		fmt.Println(string(char))
	}

}

func numJewelsInStones(J string, S string) int {
	js := make(map[rune]bool)
	for _, c := range []rune(J) {
		js[c] = true
	}
	res := 0
	for _, s := range []rune(S) {
		if js[s] {
			res++
		}
	}

	return res
}
