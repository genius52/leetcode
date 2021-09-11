package string_issue

//Input: s = "abc", k = 2
//Output: 1
//Explanation: You can split the string into "ab" and "c", and change 1 character in "ab" to make it palindrome.
func cost_palindrome(s string) int {
	begin,end := 0,len(s) - 1
	cnt := 0
	for begin < end{
		if s[begin] != s[end]{
			cnt++
		}
		begin++
		end--
	}
	return cnt
}

func dp_palindromePartition(s string,l int,begin int,k int,memo *[][]int )int{
	if begin >= l{
		return 0
	}
	//if cur_pos >= l{
	//	return cost_palindrome(s[begin:l])
	//}
	if k == 1{
		return cost_palindrome(s[begin:l])
	}
	if  (l - begin) < k{
		return 0
	}
	if (*memo)[begin][k] != -1{
		return (*memo)[begin][k]
	}

	var res int = 2147483647
	for i := begin + 1;i < l;i++{
		res = min_int(res,cost_palindrome(s[begin:i]) + dp_palindromePartition(s,l,i,k - 1,memo))
	}
	(*memo)[begin][k] = res
	return res
}

func PalindromePartition(s string, k int) int {
	var l int = len(s)
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,k + 1)
		for j := 0;j <= k;j++{
			memo[i][j] = -1
		}
	}
	return dp_palindromePartition(s,l,0,k,&memo)
}