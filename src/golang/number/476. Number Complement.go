package number

//Input: num = 5
//Output: 2
//Explanation: The binary representation of 5 is 101 (no leading zero bits),
//and its complement is 010. So you need to output 2.
func FindComplement(num int) int {
	var mask int = ^0
	for mask & num != 0{
		mask = mask << 1
	}
	return ^mask & ^num
}