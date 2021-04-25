package tree

import "strconv"

//Input: Binary tree: [1,2,3,4]
//       1
//     /   \
//    2     3
//   /
//  4
//
//Output: "1(2(4))(3)"
func tree2str(t *TreeNode) string {
	if t == nil{
		return ""
	}
	var res string = strconv.Itoa(t.Val)
	left := tree2str(t.Left)
	right := tree2str(t.Right)
	if len(left) != 0{
		res += "(" + left + ")"
	}
	if len(right) != 0{
		if len(left) == 0{
			res += "()"
		}
		res += "(" + right + ")"
	}
	return res
}