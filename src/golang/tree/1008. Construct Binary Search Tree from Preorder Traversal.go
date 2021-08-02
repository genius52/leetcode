package tree

func recursive_bstFromPreorder(preorder []int,begin int,end int)*TreeNode{
	if begin > end{
		return nil
	}
	var root TreeNode
	root.Val = preorder[begin]
	var visit int = begin + 1
	for visit <= end{
		if preorder[visit] > preorder[begin]{
			break
		}
		visit++
	}
	left := recursive_bstFromPreorder(preorder,begin + 1,visit - 1)
	right := recursive_bstFromPreorder(preorder,visit,end)
	root.Left = left
	root.Right = right
	return &root
}

func BstFromPreorder(preorder []int) *TreeNode {
	return recursive_bstFromPreorder(preorder,0,len(preorder) - 1)
}

//No recursive
//type Stack struct {
//	list *list.List
//}
//
//func NewStack() *Stack {
//	list := list.New()
//	return &Stack{list}
//}
//
//func (stack *Stack) Push(value interface{}) {
//	stack.list.PushBack(value)
//}
//
//func (stack *Stack) Pop() interface{} {
//	e := stack.list.Back()
//	if e != nil {
//		stack.list.Remove(e)
//		return e.Value
//	}
//	return nil
//}
//
//func (stack *Stack) Peak() interface{} {
//	e := stack.list.Back()
//	if e != nil {
//		return e.Value
//	}
//
//	return nil
//}
//
//func (stack *Stack) Len() int {
//	return stack.list.Len()
//}
//
//func (stack *Stack) Empty() bool {
//	return stack.list.Len() == 0
//}
//
//func bstFromPreorder(preorder []int) *TreeNode {
//	var root *TreeNode = nil
//	l := len(preorder)
//	if l == 0{
//		return root
//	}
//
//	s := NewStack()
//	var index int = 0
//	root = new(TreeNode)
//	root.Val = preorder[index]
//	s.Push(root)
//	index++
//	for index < l{
//		top := s.Peak().(*TreeNode)
//		if preorder[index] < top.Val {
//			node := new(TreeNode)
//			node.Val = preorder[index]
//			s.Peak().(*TreeNode).Left = node
//			s.Push(node)
//			top.Left = node
//			index++
//		}else{
//			var prev *TreeNode = nil
//			for !s.Empty() && preorder[index] > s.Peak().(*TreeNode).Val{
//				val,ok := (s.Pop()).(*TreeNode);if ok{
//					prev = val
//				}
//			}
//			node := new(TreeNode)
//			node.Val = preorder[index]
//			if nil != prev{
//				prev.Right = node
//			}
//			s.Push(node)
//			index++
//		}
//	}
//	return root
//}