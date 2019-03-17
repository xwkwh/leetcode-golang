package main

import "fmt"

func main() {
	points := [][]int{
		[]int{3, 3}, []int{5, -1}, []int{-2, 4},
	}
	k := 2

	fmt.Println(kClosest(points, k))
}

func kClosest(points [][]int, K int) [][]int {
	// res := make([]int, len(points))
	for i := K; i < len(points); i++ {
		for j := 0; j < K; j++ {
			s := getLen(points[i])
			ss := getLen(points[j])
			if s < ss {
				tmp := points[j]
				points[j] = points[i]
				points[i] = tmp
			}
		}
	}
	return points[:K]
}

func getLen(in []int) int {
	return in[0]*in[0] + in[1]*in[1]
}
