package diagram

//func dfs_letter(collection []string,cur_pos int,cur_str string,target int,records *[]string){
//	if cur_pos >= target{
//		if len(cur_str) >= target{
//			*records = append(*records, cur_str)
//		}
//		return
//	}
//
//	for i := 0;i < len(collection[cur_pos]);i++{
//		cur_str += string(collection[cur_pos][i])
//		dfs_letter(collection,cur_pos+1,cur_str,target,records)
//		cur_str = cur_str[:len(cur_str)-1]
//	}
//}
//
//
//func LetterCombinations(digits string) []string {
//	if len(digits) == 0{
//		return []string{}
//	}
//	var records map[int]string = make(map[int]string)
//	records[2] = "abc"
//	records[3] = "def"
//	records[4] = "ghi"
//	records[5] = "jkl"
//	records[6] = "mno"
//	records[7] = "pqrs"
//	records[8] = "tuv"
//	records[9] = "wxyz"
//
//	var collection []string
//	for _,v := range digits{
//		num,err := strconv.Atoi(string(v))
//		if err == nil{
//			collection = append(collection,records[num])
//		}
//	}
//	var res []string
//	dfs_letter(collection,0,"",len(collection),&res)
//	return res
//}

func dfs_letterCombinations(digits string,l int,idx int,record [8]string,s string,res *[]string){
	if idx >= l{
		*res = append(*res,s)
		return
	}
	pos := digits[idx] - '0' - 2
	for i := 0;i < len(record[pos]);i++{
		dfs_letterCombinations(digits,l,idx + 1,record,s + string(record[pos][i]),res)
	}
}

func LetterCombinations(digits string) []string {
	var l int = len(digits)
	if l == 0{
		return []string{}
	}
	var record [8]string
	record[0] = "abc"
	record[1] = "def"
	record[2] = "ghi"
	record[3] = "jkl"
	record[4] = "mno"
	record[5] = "pqrs"
	record[6] = "tuv"
	record[7] = "wxyz"
	var res []string
	dfs_letterCombinations(digits,l,0,record,"",&res)
	return res
}