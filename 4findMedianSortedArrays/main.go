package main

import (
	"fmt"
)

func main()  {
	i := []int{1,3}
	k := []int{2}
	res := findMedianSortedArrays(i,k)
	fmt.Println(res)
}
