package number

import "strconv"

//整数 x - 表示本回合新获得分数 x
//"+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
//"D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
//"C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
func calPoints(ops []string) int {
	var l int = len(ops)
	var record []int
	for i := 0;i < l;i++{
		if ops[i] == "+"{
			cur := record[len(record) - 1] + record[len(record) - 2]
			record = append(record,cur)
		}else if ops[i] == "D"{
			cur := record[len(record) - 1] * 2
			record = append(record,cur)
		}else if ops[i] == "C"{
			record = record[:len(record) - 1]
		}else{
			val,_ := strconv.Atoi(ops[i])
			record = append(record,val)
		}
	}
	var res int = 0
	for _,v := range record{
		res += v
	}
	return res
}