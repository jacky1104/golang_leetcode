package leetcode

import (
	"fmt"
	"testing"
)

func TestPrimeCount(t *testing.T) {

	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(200))
	fmt.Println(countPrimes(300))
}
func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if IsPrime(i) {
			count++
		}
	}

	return count

}

func IsPrime(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
