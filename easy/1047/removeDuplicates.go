package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	if len(S) == 0 {
		return ""
	}

	res := S[:1]
	l := len(S)
	for i := 0; i < l-1; i++ {
		if len(res) != 0 && S[i+1] == res[len(res)-1] {
			res = res[:len(res)-1]
			continue
		}
		res = fmt.Sprintf("%s%c", res, S[i+1])

	}

	return res

}
