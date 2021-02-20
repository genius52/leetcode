package diagram

import "strconv"

//17
func dfs_letter(collection []string,cur_pos int,cur_str string,target int,records *[]string){
	if cur_pos >= target{
		if len(cur_str) >= target{
			*records = append(*records, cur_str)
		}
		return
	}

	for i := 0;i < len(collection[cur_pos]);i++{
		cur_str += string(collection[cur_pos][i])
		dfs_letter(collection,cur_pos+1,cur_str,target,records)
		cur_str = cur_str[:len(cur_str)-1]
	}
}


func LetterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	var records map[int]string = make(map[int]string)
	records[2] = "abc"
	records[3] = "def"
	records[4] = "ghi"
	records[5] = "jkl"
	records[6] = "mno"
	records[7] = "pqrs"
	records[8] = "tuv"
	records[9] = "wxyz"

	var collection []string
	for _,v := range digits{
		num,err := strconv.Atoi(string(v))
		if err == nil{
			collection = append(collection,records[num])
		}
	}
	var res []string
	dfs_letter(collection,0,"",len(collection),&res)
	return res
}
