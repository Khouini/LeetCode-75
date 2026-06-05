package main

import "fmt"

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	rangeStart := t - 3000
	for len(this.requests) > 0 && this.requests[0] < rangeStart {
		this.requests = this.requests[1:]
	}

	return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
func main() {
	recentCounter := Constructor()

	fmt.Println(recentCounter.Ping(1))
	fmt.Println(recentCounter.Ping(100))
	fmt.Println(recentCounter.Ping(3001))
	fmt.Println(recentCounter.Ping(3002))
}
