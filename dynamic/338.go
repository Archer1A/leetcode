package main

import "fmt"

func countBit(n int) []int  {

	r :=make([]int,n+1)
	for i := 1 ; i <=n; i++{
		if i % 2 == 0{
			r[i] = r[i/2]
		}else{
			r[i] = r[i-1] + 1
		}
	}
	return r
}

func main()  {
	fmt.Println(countBit(1))
}