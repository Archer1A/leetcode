package main

import "fmt"

func generate( n int) [][]int {
	var res [][]int
	for i := 1 ; i <= n; i++ {
		if i == 1 {
			k := []int{1}
			res = append(res, k)
		}else if i == 2 {
			k := []int{1,1}
			res = append(res, k)
		}else {
			var k []int
			for j := 0 ; j < i;j ++{

				if j == 0 || j == i-1 {
					k = append(k, 1)
					continue
				}
				k = append(k, res[i-2][j-1] + res[i-2][j])
			}
			res = append(res, k)
		}
	}
    return res
}

func main()  {
	fmt.Println(generate(9))
}