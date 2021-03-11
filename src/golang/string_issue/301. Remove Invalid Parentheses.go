package string_issue


//301
//Input: "(a)())()"
//Output: ["(a)()()", "(a())()"]
func dfs_removeInvalidParentheses(s string,cur_pos int,left_tags int,right_tags int,record map[string]bool){
	if cur_pos >= len(s){
		if left_tags == right_tags{
			record[s] = true
		}
		return
	}
	if s[cur_pos] != '(' && s[cur_pos] != ')'{
		dfs_removeInvalidParentheses(s,cur_pos + 1,left_tags,right_tags,record)
	}else if s[cur_pos] == '('{
		//no action
		dfs_removeInvalidParentheses(s,cur_pos + 1,left_tags + 1,right_tags,record)
		//delete (
		s2 := s[:cur_pos] + s[cur_pos+1:]
		dfs_removeInvalidParentheses(s2,cur_pos,left_tags,right_tags,record)
	}else if s[cur_pos] == ')'{
		if right_tags == left_tags{
			//delete )
			s2 := s[:cur_pos] + s[cur_pos+1:]
			dfs_removeInvalidParentheses(s2,cur_pos,left_tags,right_tags,record)
		}else{
			//no action
			dfs_removeInvalidParentheses(s,cur_pos + 1,left_tags,right_tags + 1,record)
			//delete )
			s2 := s[:cur_pos] + s[cur_pos+1:]
			dfs_removeInvalidParentheses(s2,cur_pos,left_tags,right_tags,record)
		}
	}
}

func RemoveInvalidParentheses(s string) []string {
	var record map[string]bool = make(map[string]bool)
	dfs_removeInvalidParentheses(s,0,0,0,record)
	var max_len int = 0
	for r,_ := range record{
		if len(r) > max_len{
			max_len = len(r)
		}
	}
	var res []string
	for r,_ := range record{
		if len(r) == max_len{
			res = append(res, r)
		}
	}
	return res
}