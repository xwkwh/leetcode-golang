package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	// N: [0, N)
	// B: blacklist
	// B1: < N
	// B2: >= N
	// M: N - B1
	M   int
	mem map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	sol := Solution{
		mem: make(map[int]int),
	}
	for _, v := range blacklist {
		sol.mem[v] = -1
	}
	sol.M = N - len(sol.mem)
	for _, v := range blacklist {
		if v < sol.M {
			for contain(sol.mem, N-1) {
				N--
			}
			sol.mem[v] = N - 1
			N--
		}
	}
	return sol
}

func contain(m map[int]int, i int) bool {
	_, ok := m[i]
	return ok
	// if _, ok := m[i]; ok {
	// 	return true
	// }
	// return false
}

func (this *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(this.M)

	if contain(this.mem, x) {
		return this.mem[x]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
