package main

import "fmt"

func maxProfit(prices []int) int  {
	max := -100000000
	k := prices[0]
	if len(prices) <= 1 {
		return 0
	}
	for i := 1 ; i < len(prices) ; i ++ {
		res := prices[i] - k
		if res > max {
			max =res
		}
		// 记录最小值
		if prices[i] < k {
			k = prices[i]
		}
	}
	return max
}

func main()  {
	fmt.Println( 8&1)
}
