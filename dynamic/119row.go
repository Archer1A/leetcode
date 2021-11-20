package main

import "fmt"

func row(row int) []int {
	row += 1
	if row == 1 {
		return []int{1}
	}
	pre := []int{1,1}
	for i := 3 ; i <=row; i ++{
		var res []int
		for j := 0 ;j <i ;j ++{
			if j == 0 || j == i-1{
				res = append(res, 1)
				continue
			}
			res = append(res, pre[j]+pre[j-1])
		}
		pre = res
	}
	return pre
}
func main()  {
	fmt.Println(row(4))
}
