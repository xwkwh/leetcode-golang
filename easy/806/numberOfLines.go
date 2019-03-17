package main

import "fmt"

func main() {
	w := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"

	fmt.Println(numberOfLines(w, s))
}

func numberOfLines(widths []int, S string) []int {
	res, cur := 1, 0
	for _, v := range S {
		w := widths[v-'a']
		if cur+w > 100 {
			res += 1
			cur = w
		} else {
			cur += w
		}
	}
	return []int{res, cur}
}
