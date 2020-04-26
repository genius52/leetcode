package stack
import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}


//1190
func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		tmp := s1[i]
		s1[i] = s1[len(s1)-1-i]
		s1[len(s1)-1-i] = tmp
	}
	return string(s1)
}

func reverseParentheses(s string) string {
	stack := NewStack()
	for i := 0;i < len(s);i++{
		if s[i] == ')'{
			var sub string = ""
			for !stack.Empty(){
				v := stack.Pop().(byte)
				if v != '('{
					sub += string(v)
				}else{
					for j := 0;j < len(sub);j++{
						stack.Push(sub[j])
					}
					break
				}
			}
		}else{
			stack.Push(s[i])
		}
	}
	var res string
	for !stack.Empty(){
		res += string(stack.Pop().(byte))
	}
	return reverse(res)
}