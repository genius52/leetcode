package string_issue

//Given a list of words, each word consists of English lowercase letters.
//
//Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".
//
//A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.
//
//Return the longest possible length of a word chain with words chosen from the given list of words.
//Input: words = ["a","b","ba","bca","bda","bdca"]
//Output: 4
//Explanation: One of the longest word chain is "a","ba","bda","bdca".
func is_predecessor(s1 string,s2 string)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	var i int = 0
	var j int = 0
	var diff_cnt int = 0
	for i < l1 && j < l2{
		if s1[i] != s2[j]{
			diff_cnt++
			if diff_cnt >= 2{
				return false
			}
			j++
		}else{
			i++
			j++
		}
	}
	return true
}

func dp_longestStrChain(record [17][]string,pre_string string,cur_len int)int{
	if cur_len > 16{
		return 0
	}
	var cnt int = len(record[cur_len])
	if cnt == 0{
		return 0
	}
	//if cnt,ok := memo[pre_string];ok{
	//	return cnt
	//}
	var max_len int = 0
	for i := 0;i < cnt;i++{
		if len(pre_string) == 0 || is_predecessor(pre_string,record[cur_len][i]){
			var res int = 1 + dp_longestStrChain(record,record[cur_len][i],cur_len + 1)
			if res > max_len{
				max_len = res
			}
		}
	}
	//memo[pre_string] = max_len
	return max_len
}

func LongestStrChain(words []string) int {
	var l int = len(words)
	var record [17][]string
	for i := 0;i < l;i++{
		var cur_len int = len(words[i])
		record[cur_len] = append(record[cur_len],words[i])
	}
	var res int = 0
	for i := 1;i <= 16;i++{
		if len(record[i]) > 0{
			var cur int =  dp_longestStrChain(record,"",i)
			if cur > res{
				res = cur
			}
		}
	}
	return res
}