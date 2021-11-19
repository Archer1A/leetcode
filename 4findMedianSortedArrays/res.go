package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays (nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	 l := len(nums1)
	 if l % 2 == 1{
		 return float64(nums1[l/2])
	 }
	 sub := l / 2
	 return  float64(nums1[sub-1] + nums1[sub])/2
}

func main()  {
	i := []int{1,3}
	k := []int{2}
	res := findMedianSortedArrays(i,k)
	fmt.Println(res)
}

