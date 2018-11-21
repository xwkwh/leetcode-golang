package main

import "strings"

func main() {

}

func toLowerCase(str string) string {
	return strings.ToLower(str)
}

func toLowerCase2(wstr string) string {

	for _, c := range []rune(str) {
		if c > rune("A") {
		}
	}
	return strings.ToLower(str)
}
