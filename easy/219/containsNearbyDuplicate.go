package main

func main(){
	fmt.Println(containsNearbyDuplicate())

}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int,0)
	var has bool
	for i,v := range nums {
		if exist,ok := m[v]; ok {
			has = true
			if i-exist > k {
				return false
			}
		}

		m[v] = k
	}
	return has
} 2
