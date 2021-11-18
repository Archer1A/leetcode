package count_primenum

import (
	"fmt"
	"testing"
)

func Test_prime(t *testing.T)  {
	r := 50000
	fmt.Println(prime(r))
}

func Test_countPrimeNum(t *testing.T) {
	i := 50000
	fmt.Println( Count_PrimeNum(i))
}
func Test_values(t *testing.T)  {
	print(values())

}