package string_issue

func secondHighest(s string) int {
	var first int = -1
	var second int = -1
	for _,c := range s{
		if c >= '0' && c <= '9'{
			if first == -1{
				first = int(c)
			}else{
				if first < int(c){
					second = first
					first = int(c)
				}else if first != int(c){
					if second < int(c){
						second = int(c)
					}
				}
			}
		}
	}
	if second == -1{
		return -1
	}
	return second - '0'
}