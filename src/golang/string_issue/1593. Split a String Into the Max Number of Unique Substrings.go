package string_issue

func dp_maxUniqueSplit(s string,pos int,record map[string]bool)(bool,int){
	var l int = len(s)
	if pos >= l{
		return true,len(record)
	}
	var dup bool = false
	var max_len int = 0
	for i := pos + 1;i <= l;i++{
		sub := s[pos:i]
		if _,ok := record[sub];ok{
			continue
		}
		record[sub] = true
		ok,cur_len := dp_maxUniqueSplit(s,i,record)
		if ok{
			if cur_len > max_len{
				max_len = cur_len
			}
			dup = true
		}
		delete(record,sub)
	}
	if !dup{
		return false,max_len
	}
	return true,max_len
}

func MaxUniqueSplit(s string) int {
	//var l int = len(s)
	var record map[string]bool = make(map[string]bool)
	_,l := dp_maxUniqueSplit(s,0,record)
	return l
}