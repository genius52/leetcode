package string_issue

import "container/list"

func isValid20(s string) bool {
	var q list.List
	for _,c := range s{
		if q.Len() == 0{
			q.PushBack(c)
		}else{
			if c == '(' || c == '[' || c == '{'{
				q.PushBack(c)
			}else{
				if q.Len() == 0{
					return false
				}
				last := q.Back().Value.(int32)
				q.Remove(q.Back())
				if c == ')'{
					if last != '('{
						return false
					}
				}else if c == ']'{
					if last != '['{
						return false
					}
				}else if c == '}'{
					if last != '{'{
						return false
					}
				}
			}
		}
	}
	return q.Len() == 0
}