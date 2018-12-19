package main

import "fmt"

func main() {

	s := Constructor()
	s.Ping(1)
	s.Ping(2)
	s.Ping(3000)
	fmt.Println(s.Ping(3001))
	fmt.Println(s.Ping(3002))
}

type RecentCounter struct {
	List []int
}

func Constructor() RecentCounter {
	return RecentCounter{List: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.List = append(this.List, t)
	for this.List[0] < t-3000 {
		this.List = this.List[1:]
	}
	return len(this.List)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
