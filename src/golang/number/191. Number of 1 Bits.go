package number

//Input: n = 00000000000000000000000000001011
//Output: 3
//Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
func hammingWeight(num uint32) int {
	var res int = 0
	for i := 0;i < 32;i++{
		if (num & 1) == 1{
			res++
		}
		num = num >> 1
		if num == 0{
			break
		}
	}
	return res
}

func hammingWeight2(num uint32) int{
	var res int = 0
	for num > 0{
		if (num | 1) == num{
			res++
		}
		num >>= 1
	}
	return res
}