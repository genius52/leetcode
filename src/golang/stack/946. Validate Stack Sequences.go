package stack

import "container/list"

//Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//Output: true
//Explanation: We might do the following sequence:
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//Example 2:
//
//Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//Output: false
//Explanation: 1 cannot be popped before 2.
func ValidateStackSequences(pushed []int, popped []int) bool {
	var l int = len(pushed)
	if l == 0{
		return true
	}
	var stack list.List
	var pos_push int = 0
	var pos_pop int = 0
	for pos_push < l && pos_push < l{
		if pushed[pos_push] == popped[pos_pop]{
			pos_push++
			pos_pop++
		}else{
			if stack.Len() > 0{
				top_val := stack.Back().Value.(int)
				if top_val == popped[pos_pop]{
					stack.Remove(stack.Back())
					pos_pop++
				}else{
					stack.PushBack(pushed[pos_push])
					pos_push++
				}
			}else{
				stack.PushBack(pushed[pos_push])
				pos_push++
			}
		}
	}
	if pos_push >= l && stack.Len() > 0{
		for pos_pop < l && stack.Len() > 0{
			top_val := stack.Back().Value.(int)
			if top_val == popped[pos_pop]{
				stack.Remove(stack.Back())
				pos_pop++
			}else{
				return false
			}
		}
	}
	return pos_push == pos_pop
}