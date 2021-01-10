package string_issue

import "sort"

//Input: n = 3, k = 9
//Output: "cab"
//Explanation: There are 12 different happy string of length 3
//["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"].
//You will find the 9th string = "cab"
func visit_getHappyString(n int,cur string,record *[]string){
	if cur == ""{
		visit_getHappyString(n - 1,"a",record)
		visit_getHappyString(n - 1,"b",record)
		visit_getHappyString(n - 1,"c",record)
		return
	}
	if n == 0{
		*record = append(*record,cur)
		return
	}
	last := cur[len(cur) - 1]
	if last == 'a'{
		visit_getHappyString(n - 1,cur + "b",record)
		visit_getHappyString(n - 1,cur + "c",record)
	}else if last == 'b'{
		visit_getHappyString(n - 1,cur + "a",record)
		visit_getHappyString(n - 1,cur + "c",record)
	}else if last == 'c'{
		visit_getHappyString(n - 1,cur + "a",record)
		visit_getHappyString(n - 1,cur + "b",record)
	}
}

func GetHappyString(n int, k int) string {
	var record []string
	visit_getHappyString(n,"",&record)
	sort.Strings(record)
	if k - 1 >= len(record){
		return ""
	}
	return record[k - 1]
}