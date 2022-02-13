package string_issue

import (
	"container/list"
	"strings"
)

//返回 s 中长度为 k 且 字典序最小 的子序列，该子序列同时应满足字母 letter 出现 至少 repetition 次。
//生成的测试用例满足 letter 在 s 中出现 至少 repetition 次。
func SmallestSubsequence2(s string, k int, letter byte, repetition int) string {
	var l int = len(s)
	var rest_letter int = 0
	for i := 0;i < l;i++{
		if s[i] == letter{
			rest_letter++
		}
	}
	var q list.List
	var letter_in_q_cnt int = 0
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(s[i])
		}else{
			//出栈的条件
			//1. s[i] < q.back()
			//2. q.len + l - i > k
			//3. total > repetition
			var replace bool = false
			for q.Len() > 0 && q.Back().Value.(byte) > s[i] && letter_in_q_cnt + rest_letter >= repetition && q.Len() + l - i >= k{
				if q.Back().Value.(byte) == letter{
					letter_in_q_cnt--
					if letter_in_q_cnt + rest_letter < repetition{
						letter_in_q_cnt++
						break
					}
				}
				replace = true
				q.Remove(q.Back())
			}
			if replace || q.Len() < k || (s[i] == letter && letter_in_q_cnt + rest_letter == repetition){
				q.PushBack(s[i])
				if s[i] == letter{
					letter_in_q_cnt++
				}
			}
		}
		if s[i] == letter{
			rest_letter--
		}
	}
	var data []byte
	for q.Len() > 0{
		data = append(data,q.Front().Value.(byte))
		q.Remove(q.Front())
	}
	idx := len(data) - 1
	for len(data) > k && idx >= 0{
		for data[idx] == letter{
			idx--
		}
		data = append(data[:idx],data[idx + 1:]...)
		idx--
	}
	var sb strings.Builder
	for i := 0;i < len(data);i++{
		sb.WriteByte(data[i])
	}
	return sb.String()
}