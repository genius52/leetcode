package number

//Input: "2*3-4*5"
//Output: [-34, -14, -10, -10, 10]
//Explanation:
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10
//func dp_diffWaysToCompute(input string,last_sum int,last_opt uint8,record *[]int){
//	l := len(input)
//	if l == 0{
//		*record = append(*record,last_sum)
//		return
//	}
//	i := 0
//	begin := 0
//	for i < (l - 1) && input[i + 1] != '+' && input[i + 1] != '-' && input[i + 1] != '*'{
//		i++;
//	}
//	n,err := strconv.ParseInt(input[i:begin + 1],10,32)
//	if err != nil{
//		return
//	}
//	tmp := last_sum
//	switch last_opt {
//	case '+':
//		last_sum += int(n)
//	case '-':
//		last_sum -= int(n)
//	case '*':
//		last_sum *= int(n)
//	case 0:
//		last_sum = int(n)
//	}
//	dp_diffWaysToCompute(input[i + 2:],last_sum,input[i + 1],record)
//	return dp_diffWaysToCompute(input[i + 2:],int(n),input[i + 1],record)
//}
//
//func diffWaysToCompute(input string) []int {
//	var res []int
//	dp_diffWaysToCompute(input,0,0,&res)
//	return res
//}
