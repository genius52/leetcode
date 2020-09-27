package array

func minOperations(logs []string) int {
	var l int = len(logs)
	var depth int = 0
	for i := 0;i < l;i++{
		if logs[i] == "../"{
			if depth > 0{
				depth--
			}
		}else{
			if logs[i] != "./"{
				depth++
			}
		}
	}
	return depth
}
