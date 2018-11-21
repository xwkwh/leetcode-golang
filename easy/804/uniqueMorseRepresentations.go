package main

import "fmt"

func main() {
	fmt.Println(rune('a'))
	fmt.Println('b'-'a')
	fmt.Println(rune('c')-'a')

	words := []string{"gin", "zen", "gig", "msg"}

	fmt.Println(uniqueMorseRepresentations(words))
}

var morse = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}


func uniqueMorseRepresentations(words []string) int {
	ret := 0
	memory := make(map[string]bool)
	for _, word := range words {
		var tmp string
		for _, c := range  []rune(word) {
			tmp += morse[c - 'a']
		}
		fmt.Println(tmp)
		if !memory[tmp] {
			memory[tmp] = true
			ret ++
		}
	}
	return  ret
}