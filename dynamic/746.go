package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	c := make([]int,len(cost))
	for curr := 2 ; curr< len(cost) ; curr ++{
		a := c[curr-1] + cost[curr-1]
		b := c[curr-2] +cost[curr-2]
		if a <= b {
			c[curr] +=  a
		}else {
			c[curr] +=  b
		}
	}
	return c[len(cost)-1]
}

func main()  {
	list := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(list))
}