package string_issue

//Input: word1 = "cabaa", word2 = "bcaaa"
//Output: "cbcabaaaaa"
//"uuurr"
//"urrru"
func LargestMerge(word1 string, word2 string) string {
	var l1 int = len(word1)
	var l2 int = len(word2)
	if l1 == 0 && l2 == 0{
		return ""
	}
	if l2 == 0 {
		return word1
	}
	if l1 == 0 {
		return word2
	}
	if word1 > word2{
		return string(word1[0]) + LargestMerge(word1[1:],word2)
	}else{
		return string(word2[0]) + LargestMerge(word1,word2[1:])
	}
}