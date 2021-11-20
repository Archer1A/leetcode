package main

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	sum := 0
	for _, i := range nums{
		sum += i
		res = int(math.Max(float64(res),float64(sum)))
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

func main()  {
	i := []int{-2,1,-3,4,-1,2,1,-5,4}
	print(maxSubArray(i))
}
