package diagram

import "math"

//req_skills = ["algorithms","math","java","reactjs","csharp","aws"],

//people = [["algorithms","math","java"],["algorithms","math","reactjs"],
//["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]

var cur_min int = 2147483647
func dp_smallestSufficientTeam(people_skills []int,idx int,cur_state int,target int,choose *[]int,res *[][]int){
	if cur_state == target{
		cur_len := len(*choose)
		if cur_len < cur_min{
			cur_min = cur_len
		}else{
			return
		}
		var cur []int = make([]int,cur_len)
		copy(cur,*choose)
		*res = append(*res,cur)
	}
	var people_len int = len(people_skills)
	if idx >= people_len{
		return
	}
	if len(*choose) >= (cur_min - 1){
		return
	}
	//choose current people
	//var cur_choose []int = make([]int,len(*choose))

	//copy(cur_choose,*choose)
	var next_idx int = idx
	for next_idx < people_len && (cur_state | people_skills[next_idx]) == cur_state{
		next_idx++
	}
	if next_idx < people_len{
		*choose = append(*choose,next_idx)
		dp_smallestSufficientTeam(people_skills,next_idx + 1,cur_state | people_skills[next_idx],target,choose,res)
		*choose = (*choose)[0:len(*choose) - 1]
	}
	//skip current people
	dp_smallestSufficientTeam(people_skills,next_idx + 1,cur_state,target,choose,res)
}

func SmallestSufficientTeam(req_skills []string, people [][]string) []int {
	cur_min = len(people)
	var l int = len(req_skills)
	var target int = int(math.Pow(2,float64(l))) - 1
	var skill_bit map[string]int = make(map[string]int)
	for i := 0;i < l;i++{
		skill_bit[req_skills[i]] = i
		target = target | (1 << i)
	}
	var people_skills []int = make([]int,len(people))
	for i := 0;i < len(people);i++{
		var cur_len int = len(people[i])
		var cur int = 0
		for j := 0;j < cur_len;j++{
			if val,ok := skill_bit[people[i][j]];ok{
				cur = cur | (1 << val)
			}
		}
		people_skills[i] = cur
	}
	var res [][]int
	dp_smallestSufficientTeam(people_skills,0,0,target,&[]int{},&res)
	var min_len int = 2147483647
	var min_idx int = -1
	for i := 0;i < len(res);i++{
		if len(res[i]) < min_len{
			min_len = len(res[i])
			min_idx = i
		}
	}
	return res[min_idx]
}