package number

import "strconv"

//cost.length == 9
//1 <= cost[i] <= 5000
//1 <= target <= 5000

//True: s1 > s2
func my_cmp_str(s1 string,s2 string)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 > l2{
		return true
	}else if l1 < l2{
		return false
	}
	for i := 0;i < l1;i++{
		if s1[i] > s2[i]{
			return true
		}else if s1[i] < s2[i]{
			return false
		}else{
			continue
		}
	}
	return true
}

func dfs_largestNumber(cost []int,target int,memo *[]string)(string){
	if target == 0{
		return ""
	}
	if target < 0{
		return "0"
	}
	if len((*memo)[target]) > 0{
		return (*memo)[target]
	}
	var res string = "0"
	for i := 8;i >= 0;i--{
		rest_target := target - cost[i]
		if rest_target < 0{
			continue
		}
		next_str := dfs_largestNumber(cost,rest_target,memo)
		if next_str == "0"{
			continue
		}
		s2 := strconv.Itoa(i + 1) + next_str
		if my_cmp_str(s2,res){
			res = s2
		}
	}
	(*memo)[target] = res
	return res
}

func LargestNumber2(cost []int, target int) string {
	var memo []string = make([]string,target + 1)
	return dfs_largestNumber(cost,target,&memo)
}