package main

func main() {

}

func transpose(A [][]int) [][]int {
	m := len(A)
	if m == 0 {
		return nil
	}
	n := len(A[0])
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, m)
	}
	for k, v := range A {
		for kk, vv := range v {
			b[kk][k] = vv
		}
	}
	return b

}
