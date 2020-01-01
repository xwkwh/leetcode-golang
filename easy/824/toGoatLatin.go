package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("Each word consists of lowercase and uppercase letters only"))

}

func toGoatLatin(S string) string {
	mem := map[rune]bool{
		'a':true,
		'e':true,
		'i':true,
		'o':true,
		'u':true,
	}

	ss := strings.Split(S," ")
	for k,v := range ss {
		vv := strings.ToLower(v)
		if mem[rune(vv[0])] {
			tmp := fmt.Sprintf("%sma",v)
			for i := 0 ;i<= k;i ++ {
				tmp = fmt.Sprintf("%sa",tmp)
			}
			ss[k] = tmp
			continue
		}

		tmp := fmt.Sprintf("%s%sma",v[1:],v[:1])
		for i := 0 ;i<= k;i ++ {
			tmp = fmt.Sprintf("%sa",tmp)
		}
		ss[k] = tmp

	}

	return strings.Join(ss," ")

}
