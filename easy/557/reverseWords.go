package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "Let's take LeetCode contest"
	out := "s'teL ekat edoCteeL tsetnoc"

	fmt.Println(out == reverseWords(in))

}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for k, v := range ss {
		ss[k] = revers(v)
	}
	return strings.Join(ss, " ")
}

func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
