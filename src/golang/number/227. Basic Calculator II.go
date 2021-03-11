package number

import (
	"container/list"
	"strconv"
)

func Calculate227(s string) int {
	var q list.List
	var cur_num int = 0
	var last_op int32
	s = "+" + s + "+"
	for _,c := range s{
		if c >= '0' && c <= '9' {
			n , _ := strconv.Atoi(string(c))
			cur_num = cur_num * 10 + n
			continue
		}
		if c == ' '{
			continue
		}
		if last_op == '+'{
			q.PushBack(cur_num)
		}else if last_op == '-'{
			q.PushBack(-cur_num)
		}else if last_op == '*'{
			last_num := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushBack(last_num * cur_num)
		}else if last_op == '/'{
			last_num := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushBack(last_num/cur_num)
		}
		last_op = c
		cur_num = 0
	}
	var sum int = 0
	for item := q.Front();nil != item ;item = item.Next() {
		n := item.Value.(int)
		sum += n
	}
	return sum
}