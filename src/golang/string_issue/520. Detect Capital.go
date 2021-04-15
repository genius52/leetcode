package string_issue

func detectCapitalUse(word string) bool {
	var l int = len(word)
	if l <= 1{
		return true
	}
	var first_is_capital bool = (word[0] >= 'A' && word[0] <= 'Z')
	var second_is_capital bool = (word[1] >= 'A' && word[1] <= 'Z')
	for i := 2;i < l;i++{
		if word[i] >= 'a' && word[i] <= 'z'{
			if(second_is_capital){
				return false;
			}
		}else{
			if(!second_is_capital){
				return false;
			}
		}
	}
	if(!first_is_capital && second_is_capital){
		return false;
	}
	return true
}