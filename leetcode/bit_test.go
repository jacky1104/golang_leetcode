package leetcode

import "fmt"

func singleNumber(nums []int) int {

	result := 0
	for _, v := range nums {
		result = result ^ v
	}
	return result

}

/**
1. a[i] , 0
2. 0, a[i]
3. 0, 0
*/
func singleNumber2(nums []int) int {

	one, two := 0, 0

	for _, v := range nums {

		one = (one ^ v) & (^two)
		two = (two ^ v) & (^one)
	}

	return one
}

func singleNumber22(nums []int) int {

	one, two, three := 0, 0, 0

	for _, v := range nums {

		two |= one & v
		one = one ^ v
		three = one & two
		one &= one & (^three)
		two &= two & (^three)
	}

	return one
}

func singleNumber3(nums []int) []int {

	diff := 0

	for _, v := range nums {
		diff = diff ^ v
	}

	// find the last different bit
	diff &= -diff

	result := []int{0, 0}

	for _, v := range nums {

		if diff&v == 0 {
			result[0] ^= v
		} else {
			result[1] ^= v
		}
	}

	return result

}

func reverseBits(num uint32) uint32 {

	var res uint32
	for i := 31; i >= 0; i-- {
		res = res | ((num & 1) << uint(i))
		num = num >> 1
	}
	return res
}

func hammingWeight(num uint32) int {

	res := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	fmt.Println(num)
	return res
}
