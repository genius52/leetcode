package number

//Input: nums = [0,1,1,1,1,1]
//Output: [true,false,false,false,true,false]
func prefixesDivBy5(A []int) []bool {
	var l int = len(A)
	var res []bool = make([]bool,l)
	var n int = 0
	for i := 0;i < l;i++{
		n = (n * 2 + A[i]) % 5
		if n == 0{
			res[i] = true
		}
	}
	return res
}