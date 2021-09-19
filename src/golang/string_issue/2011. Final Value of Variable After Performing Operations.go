package string_issue

func finalValueAfterOperations(operations []string) int {
	var l int = len(operations)
	var res int = 0
	for i := 0;i < l;i++{
		if operations[i][1] == '+'{
			res++
		}else{
			res--
		}
	}
	return res
}