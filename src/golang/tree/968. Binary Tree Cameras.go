package tree

//return 1:
//status:
//0: child not covered.
//1: child is covered
//2: child has camera
//return 2:total cameras on child branch
//func dfs_minCameraCover(node *TreeNode)(int,int) {
//	if node == nil{
//		return 1,0
//	}
//	if node.Left == nil && node.Right == nil{
//		return 0,0
//	}
//	if node.Left == nil{
//		status,total_right := dfs_minCameraCover(node.Right)
//		if status == 0{
//			return 2,1 + total_right
//		}else if status == 1 {
//			return 0, total_right
//		}
//		return 1,total_right
//	}else if node.Right == nil{
//		status,total_left := dfs_minCameraCover(node.Left)
//		if status == 0{
//			return 2,1 + total_left
//		}else if status == 1{
//			return 0,total_left
//		}
//		return 1,total_left
//	}else{
//		left_status,total_left := dfs_minCameraCover(node.Left)
//		right_status,total_right := dfs_minCameraCover(node.Right)
//		if left_status == 0 || right_status == 0{
//			return 2,1 + total_left + total_right
//		}else if (left_status == 2 && right_status != 0) || (right_status == 2 && left_status != 0) {
//			return 1, total_left + total_right
//		}
//		return 0, total_left + total_right
//	}
//}
//
//func minCameraCover(root *TreeNode) int {
//	if root == nil{
//		return 0
//	}
//	if root.Left == nil && root.Right == nil{
//		return 1
//	}
//	status,total := dfs_minCameraCover(root)
//	if status == 0{
//		return total + 1
//	}
//	return total
//}

const (
	empty int = 0
	camera int = 1
	watched int = 2
)

func dfs_minCameraCover(node *TreeNode,res *int)int {
	if node == nil{
		return empty
	}
	if node.Left == nil && node.Right == nil{
		return empty
	}
	if node.Left == nil{
		status := dfs_minCameraCover(node.Right,res)
		if status == empty{
			*res++
			return camera
		}else if status == camera {
			return watched
		}else if status == watched{
			return empty
		}
	}else if node.Right == nil{
		status := dfs_minCameraCover(node.Left,res)
		if status == empty{
			*res++
			return camera
		}else if status == camera {
			return watched
		}else if status == watched{
			return empty
		}
	}
	left_status := dfs_minCameraCover(node.Left,res)
	right_status := dfs_minCameraCover(node.Right,res)
	if left_status == empty || right_status == empty {
		*res++
		return camera
	}else{
		if left_status == camera || right_status == camera{
			return watched
		}else{
			return empty
		}
	}
}

func minCameraCover(root *TreeNode) int {
	var res int = 0
	status := dfs_minCameraCover(root,&res)
	if status == empty{
		res++
	}
	return res
}