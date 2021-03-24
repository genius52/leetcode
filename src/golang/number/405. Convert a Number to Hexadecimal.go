package number

//Input:
//26
//
//Output:
//"1a"
func ToHex(num int) string {
	if num == 0{
		return "0"
	}
	var num2 uint32 = uint32(num)
	var dict string = "0123456789abcdef"
	var res string
	for num2 > 0{
		mod := num2 % 16
		res = string(dict[mod]) + res
		num2 = (num2 >> 4)
	}
	return res
}