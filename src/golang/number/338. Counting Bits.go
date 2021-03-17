package number

//Input: num = 5
//Output: [0,1,1,2,1,2]
//Explanation:
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
func CountBits(num int) []int {
	var res []int = make([]int,num + 1)
	for i := 1;i <= num;i++{
		if (i | 1) == i{
			res[i] = res[i/2] + 1
		}else{
			res[i] = res[i/2]
		}
	}
	return res
}