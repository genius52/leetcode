package array

//Now say a word a from A is universal if for every b in B, b is a subset of a.
//Return a list of all universal words in A.  You can return the words in any order.
//Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
//Output: ["facebook","leetcode"]
func compare_wordSubsets(a string,word_record [26]int)bool{
	var a_record [26]int
	for _,c := range a{
		a_record[c - 'a']++
	}
	for i := 0;i < 26;i++{
		if word_record[i] == 0{
			continue
		}
		if a_record[i] == 0 || a_record[i] < word_record[i]{
			return false
		}
	}
	return true
}

func WordSubsets(A []string, B []string) []string {
	var res []string
	var record [26]int
	for _,b := range B{
		var word_count [26]int
		for _,c := range b{
			word_count[c - 'a']++
		}
		for i := 0;i < 26;i++{
			if word_count[i] > record[i]{
				record[i] = word_count[i]
			}
		}
	}
	for _,a := range A{
		if !compare_wordSubsets(a,record){
			continue
		}
		res = append(res,a)
	}
	return res
}