package main

import "fmt"

func isSubsequence(s string, t string) bool {

	if len(s) ==0 {
		return true
	}
	if len(t) ==0 {
		return false
	}
	var dp [][]bool
	for j,k :=  range s {
		var p []bool
		for m,n := range t{
			if k == n {
				if j != 0 && m != 0 {
					p = append(p, dp[j-1][m-1])
				}else {
					p = append(p, true)
				}
			}else  {
				//前一个为true 后面都为true
				if m !=0 {
					p = append(p, p[m-1])
				}else  {
					p = append(p, false)
				}
			}
		}
		dp = append(dp, p)
	}
	return dp[len(s) -1][len(t)-1]
}

func main()  {
	fmt.Println(isSubsequence("abc","asdacwbaa"))
}