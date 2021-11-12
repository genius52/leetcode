package string_issue

func dp_maxUniqueSplit(s string,pos int,record map[string]bool,max_len *int){
	var l int = len(s)
	if pos >= l{
		if len(record) > *max_len{
			*max_len = len(record)
		}
	}
	for i := pos + 1;i <= l;i++{
		sub := s[pos:i]
		if _,ok := record[sub];ok{
			continue
		}
		record[sub] = true
		dp_maxUniqueSplit(s,i,record,max_len)
		delete(record,sub)
	}
}

func MaxUniqueSplit(s string) int {
	var record map[string]bool = make(map[string]bool)
	var max_len int = 0
	dp_maxUniqueSplit(s,0,record,&max_len)
	return max_len
}