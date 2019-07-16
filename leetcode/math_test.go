package leetcode

/**
9. Palindrome Number
https://leetcode.com/problems/palindrome-number/
*/
func isPalindrome(x int) bool {

	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	temp := x
	reverse := 0

	for temp > 0 {
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}

	if reverse == x {
		return true
	}

	return false

}
