package number

import "strconv"

//Input: "199100199"
//Output: true
//Explanation: The additive sequence is: 1, 99, 100, 199.
//             1 + 99 = 100, 99 + 100 = 199
func dfs_isAdditiveNumber(num string,last_num int,target_num int)bool{
	l := len(num)
	if l == 0{
		return true
	}
	if l > 1 && num[0] == '0'{
		return false
	}
	for i := 1;i <= l;i++{
		n,_ := strconv.Atoi(num[:i])
		if n > target_num{
			break
		}
		if n == target_num{
			if dfs_isAdditiveNumber(num[i:],target_num,last_num + target_num){
				return true
			}
		}
	}
	return false
}

func IsAdditiveNumber(num string) bool {
	l := len(num)
	if l < 3{
		return false
	}
	for i := 1;i <= l / 2;i++{
		if num[i] == '0'{
			first_num,_ :=  strconv.Atoi(num[:i])
			if dfs_isAdditiveNumber(num[i+1:],0,first_num){
				return true
			}
			if num[0] == 0{
				break
			}
			continue
		}
		for j := i + 1;j < l ;j++{
			first_num,_ :=  strconv.Atoi(num[:i])
			second_num,_ := strconv.Atoi(num[i:j])
			if dfs_isAdditiveNumber(num[j:],second_num,first_num + second_num){
				return true
			}
		}
		if num[0] == '0'{
			break
		}
	}
	return false
}