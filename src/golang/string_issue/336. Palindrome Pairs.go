package string_issue

//isPalindrome
//Input: words = ["abcd","dcba","lls","s","sssll"]
//Output: [[0,1],[1,0],[3,2],[2,4]]
//Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
func check_valid(s1 string,s2 string)bool{//check s1 + s2
	var l1 int = len(s1)
	var l2 int = len(s2)
	var l int = l1 + l2
	//len = 4, 1  2
	//len = 3,  0,2
	var pos1 int = l/2 - 1
	var pos2 int = l/2
	if (l | 1) == l{
		pos2++
	}
	for pos1 >= 0 && pos2 < l{
		if pos1 >= l1{
			if s2[pos1 - l1] != s2[pos2 - l1]{
				return false
			}
		}else if pos2 < l1{
			if s1[pos1] != s1[pos2]{
				return false
			}
		}else{//pos1 at s1,pos2 at s2
			if s1[pos1] != s2[pos2 - l1]{
				return false
			}
		}
		pos1--
		pos2++
	}
	return true
}

func PalindromePairs(words []string) [][]int {
	var l int = len(words)
	var res [][]int
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			ret := check_valid(words[i],words[j])
			if ret{
				res = append(res,[]int{i,j})
			}
			ret = check_valid(words[j],words[i])
			if ret{
				res = append(res,[]int{j,i})
			}
		}
	}
	return res
}