package string_issue

//Input: "aab"
//Output:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
func recursive_string(s string,begin int,record *[][]string,previous []string) {
	if begin < 0 || begin >= len(s) {
		return
	}
	for i,j := begin,begin + 1;i < j && j < len(s);j++{
		if isPalindrome(s[i:j]) {
			if  isPalindrome(s[j:]) {
				var res []string
				res = append(res,previous...)
				res = append(res,s[i:j])
				res = append(res,s[j:])
				*record = append(*record,res)
				recursive_string(s,j,record,append(previous,s[i:j]))
			}else{
				recursive_string(s,j,record,append(previous,s[i:j]))
			}
		}
	}
}

func partition(s string) [][]string {
	var res [][]string
	recursive_string(s,0,&res,nil)
	if isPalindrome(s){
		res = append(res, []string{s})
	}
	return res
}