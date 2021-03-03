package number

import "strconv"

//Input: num = "232", target = 8
//Output: ["2*3+2", "2+3*2"]

//Input: num = "105", target = 5
//Output: ["1*0+5","10-5"]
func dfs_addOperators(num string,l int,pos int,exp string,pre_sum int,last_num int,target int,res *[]string){
	if pos >= l{
		if pre_sum == target{
			*res = append(*res,exp)
		}
		return
	}
	//insert +
	//insert -
	//insert *
	for i := pos;i < l;i++{
		cur_num,err := strconv.Atoi(num[pos:i+1])
		if len(strconv.Itoa(cur_num)) != (i + 1 - pos){
			continue
		}
		if err == nil{
			dfs_addOperators(num,l,i + 1,exp + "+" + num[pos:i + 1],pre_sum + cur_num,cur_num,target,res)
			dfs_addOperators(num,l,i + 1,exp + "-" + num[pos:i + 1],pre_sum - cur_num,-cur_num,target,res)
			//multiple last number
			dfs_addOperators(num,l,i + 1,exp + "*" + num[pos:i + 1],
				pre_sum - last_num + last_num * cur_num,last_num * cur_num,target,res)
		}

	}
}

func AddOperators(num string, target int) []string {
	var l int = len(num)
	var res []string
	for i := 0;i < l;i++{
		val,err := strconv.Atoi(num[0:i + 1])
		if err != nil{
			continue
		}
		if len(strconv.Itoa(val)) != (i + 1){
			continue
		}
		var exp string = strconv.Itoa(val)
		dfs_addOperators(num,l,i + 1,exp,val,val,target,&res)
	}
	return res
}