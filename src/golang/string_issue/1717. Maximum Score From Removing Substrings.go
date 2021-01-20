package string_issue

import "container/list"

//Input: s = "cdbcbbaaabab", x = 4, y = 5
//Output: 19
//Explanation:
//- Remove the "ba" underlined in "cdbcbbaaabab". Now, s = "cdbcbbaaab" and 5 points are added to the score.
//- Remove the "ab" underlined in "cdbcbbaaab". Now, s = "cdbcbbaa" and 4 points are added to the score.
//- Remove the "ba" underlined in "cdbcbbaa". Now, s = "cdbcba" and 5 points are added to the score.
//- Remove the "ba" underlined in "cdbcba". Now, s = "cdbc" and 5 points are added to the score.
//Total score = 5 + 4 + 5 + 5 = 19.
func MaximumGain(s string, x int, y int) int {
	var first uint8 = 'a'
	var second uint8 = 'b'
	var min_score,max_score int = y,x
	if x < y{
		first = 'b'
		second = 'a'
		min_score = x
		max_score = y
	}

	var l int = len(s)
	var res int = 0
//aabbaaxybbaabb
	var q list.List
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(s[i])
		}else{
			top := q.Back().Value.(uint8)
			if top == first && s[i] == second{
				res += max_score
				q.Remove(q.Back())
			}else{
				q.PushBack(s[i])
			}
		}
	}
	var q2 list.List
	for q.Len() > 0{
		if q2.Len() == 0{
			q2.PushBack(q.Front().Value.(uint8))
		}else{
			q_front := q.Front().Value.(uint8)
			q2_back := q2.Back().Value.(uint8)
			if q_front == first && q2_back == second{
				res += min_score
				q2.Remove(q2.Back())
			}else{
				q2.PushBack(q_front)
			}
		}
		q.Remove(q.Front())
	}
	return res
}