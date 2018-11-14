package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
	ss := strings.Split("b   a    ", " ")
	fmt.Println("ss", len(ss))
	fmt.Print(ss[len(ss)-1])
	fmt.Print(ss[len(ss)-1] == "")
	fmt.Println(len(ss[len(ss)-1]))
	fmt.Println(lengthOfLastWord("aa "))
	fmt.Println(lengthOfLastWord("b   a    "))
}

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	ss := strings.Split(s, " ")
	last := len(ss) - 1
	for ss[last] == "" && last > 0 {
		last = last - 1
	}
	return len(ss[last])
}
