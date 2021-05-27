package string_issue

//Input: s = "ababcbacadefegdehijhklij"
//Output: [9,7,8]
//Explanation:
//The partition is "ababcbaca", "defegde", "hijhklij".
//This is a partition so that each letter appears in at most one part.
//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
func PartitionLabels(S string) []int{
	var last_appear [26]int
	for pos,c := range S{
		last_appear[c - 'a'] = pos
	}
	var res []int
	var l int = len(S)
	var left int = 0
	var right int = left
	var pre_last int = 0
	for left < l{
		for right < l{
			if pre_last < last_appear[S[right] - 'a']{
				pre_last = last_appear[S[right] - 'a']
			}
			if right == pre_last{
				break
			}
			right++
		}
		res = append(res,right - left + 1)
		left = right + 1
		right = left
	}
	return res
}

func partitionLabels(S string) []int {
	var res []int
	var dict [26]int
	for pos,e := range S{
		dict[e - 'a'] = pos
	}
	l := len(S)
	i := 0
	start := 0
	last := 0
	for i < l {
		if dict[S[i] - 'a'] > last{
			last = dict[S[i] - 'a']
		}
		if last == i{
			if last == i{
				res = append(res, last - start + 1)
				start = i + 1
			}
		}
		i++
	}
	return res
}