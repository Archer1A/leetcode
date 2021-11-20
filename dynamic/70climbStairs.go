package main

import "fmt"

func ClimbStairs(n int) int {
	i :=make([]int,n+1)
	i[0] = 1
	for j := 1 ; j <= n; j++{
		a := 0
		b := 0
		if j - 1 >= 0 {
			a = i[j-1]
		}
		if j - 2 >= 0 {
			b = i[j-2]
		}
		i[j] = a+b
	}
	return i[n]
}

func main()  {
	fmt.Println(ClimbStairs(4))
}