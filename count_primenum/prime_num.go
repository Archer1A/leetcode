package count_primenum

var p  = map[int]int{}

func Count_PrimeNum(n int) int {
	if n  <= 1 {
		return 0
	}
	count := 0
	for i := 2 ;i <= n ;i++ {
		count += prime(i)
	}
	return count
}

func prime(num int) int {
	j := num
	count := 0
	for i := 2 ; i*i <= j ; i ++ {
		for  j % i == 0 {
			count ++
			j /=  i
		}
	}
	if j > 1 {
		count ++
	}
	return count
}

func values() (result int)  {

	defer func() {
		result ++
	}()
	return result
}