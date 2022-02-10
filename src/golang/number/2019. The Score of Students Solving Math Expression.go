package number

import "container/list"

func dfs_scoreOfStudents(s string,record map[string]map[int]bool)map[int]bool{
	var l int = len(s)
	var res map[int]bool = make(map[int]bool)
	if l == 1{
		res[int(s[0] - '0')] = true
		return res
	}
	if _,ok := record[s];ok{
		return record[s]
	}
	for i := 1;i < l;i += 2{
		left_set := dfs_scoreOfStudents(s[:i],record)
		right_set := dfs_scoreOfStudents(s[i + 1:],record)
		for k1,_ := range left_set{
			for k2,_ := range right_set{
				if s[i] == '+'{
					if k1 + k2 <= 1000{
						res[k1 + k2] = true
					}
				}else {
					if k1 * k2 <= 1000{
						res[k1 * k2] = true
					}
				}
			}
		}
	}
	record[s] = res
	return res
}

func ScoreOfStudents(s string, answers []int) int{
	var l int = len(s)
	var right_answer int = 0
	var q list.List
	q.PushBack(int(s[0] - '0'))
	for i := 1;i < l;i += 2{
		if s[i] == '+'{
			q.PushBack(int(s[i + 1] - '0'))
		}else if s[i] == '*'{
			last := q.Back().Value.(int)
			q.Remove(q.Back())
			last *= int(s[i + 1] - '0')
			q.PushBack(last)
		}
	}
	for q.Len() > 0{
		right_answer += q.Back().Value.(int)
		q.Remove(q.Back())
	}
	var record map[string]map[int]bool = make(map[string]map[int]bool)
	result := dfs_scoreOfStudents(s,record)
	var res int = 0
	for _,ans := range answers{
		if ans == right_answer{
			res += 5
		}else if _,ok := result[ans];ok{
			res += 2
		}
	}
	return res
}

//func dfs_scoreOfStudents(s string,l int,pos int,data []int,record map[int]bool){
//	if pos >= l{
//		var res int = 0
//		for i := 0;i < len(data);i++{
//			res += data[i]
//		}
//		record[res] = true
//		return
//	}
//	if s[pos] == '+'{
//		data = append(data,int(s[pos + 1] - '0'))
//		dfs_scoreOfStudents(s,l,pos + 2,data,record)
//	}else if s[pos] == '*'{
//		//multiple last number
//		var cur_len int = len(data)
//		pre := data[cur_len - 1]
//		next_num := int(s[pos + 1] - '0')
//		pre *= next_num
//		var data2 []int = make([]int,cur_len)
//		copy(data2,data)
//		data2[cur_len - 1] = pre
//		dfs_scoreOfStudents(s,l,pos + 2,data2,record)
//		//multiple all previous number with next one number
//		var res int = 0
//		for i := 0;i < cur_len;i++{
//			res += data[i]
//		}
//		dfs_scoreOfStudents(s,l,pos + 2,[]int{res * next_num},record)
//		//multiple all previous numbers with all next numbers
//	}
//}
//
//func ScoreOfStudents(s string, answers []int) int {
//	var l int = len(s)
//	var right_answer int = 0
//	var q list.List
//	q.PushBack(int(s[0] - '0'))
//	for i := 1;i < l;i += 2{
//		if s[i] == '+'{
//			q.PushBack(int(s[i + 1] - '0'))
//		}else if s[i] == '*'{
//			last := q.Back().Value.(int)
//			q.Remove(q.Back())
//			last *= int(s[i + 1] - '0')
//			q.PushBack(last)
//		}
//	}
//	for q.Len() > 0{
//		right_answer += q.Back().Value.(int)
//		q.Remove(q.Back())
//	}
//	var data []int = []int{int(s[0] - '0')}
//	var record map[int]bool = make(map[int]bool)
//	dfs_scoreOfStudents(s,l,1,data,record)
//	var res int = 0
//	for _,ans := range answers{
//		if ans == right_answer{
//			res += 5
//		}else if _,ok := record[ans];ok{
//			res += 2
//		}
//	}
//	return res
//}