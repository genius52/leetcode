package tree
//95
//dp Top to bottom with memory
func dfs_generateTrees(low int,high int,record [][][]*TreeNode)[]*TreeNode{
	if low <= 0 || low > high{
		return nil
	}
	if record[low][high] != nil{
		return record[low][high]
	}
	var root_list []*TreeNode
	for i := low;i <= high;i++{
		left_list := dfs_generateTrees(low,i - 1,record)
		right_list := dfs_generateTrees(i + 1,high,record)
		if left_list != nil && right_list != nil{
			for _,r := range right_list{
				for _,l := range left_list{
					var root TreeNode
					root.Val = i
					root.Left = l
					root.Right = r
					root_list = append(root_list,&root)
				}
			}
		}else if left_list == nil{
			for _,r := range right_list{
				var root TreeNode
				root.Val = i
				root.Left = nil
				root.Right = r
				root_list = append(root_list,&root)
			}
		}else if right_list == nil{
			for _,l := range left_list{
				var root TreeNode
				root.Val = i
				root.Left = l
				root.Right = nil
				root_list = append(root_list,&root)
			}
		}
	}
	record[low][high] = root_list
	return root_list
}

func GenerateTrees(n int) []*TreeNode {
	var record [][][]*TreeNode = make([][][]*TreeNode,n + 1)
	for i := 0;i <= n;i++{
		record[i] = make([][]*TreeNode,n + 1)
		var node TreeNode
		node.Val = i
		record[i][i] = []*TreeNode{&node}
	}
	return dfs_generateTrees(1,n,record)
}
