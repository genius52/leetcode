package tree

//Input: root1 = [2,1,4], root2 = [1,0,3]
//Output: [0,1,1,2,3,4]
func pre_visit1305(node *TreeNode,record *[]int){
	if node == nil {
		return
	}
	pre_visit1305(node.Left,record)
	*record = append(*record, node.Val)
	pre_visit1305(node.Right,record)
}

func merge_array(record1 []int,record2 []int)[]int{
	l1 := len(record1)
	l2 := len(record2)
	if l1 == 0{
		return record2
	}
	if l2 == 0{
		return record1
	}
	var res []int = make([]int,l1 + l2)
	i := 0
	h1 := 0
	h2 := 0
	for h1 < l1 && h2 < l2{
		if record1[h1] < record2[h2]{
			res[i] = record1[h1]
			h1++
		}else{
			res[i] = record2[h2]
			h2++
		}
		i++
	}
	if h1 == l1{
		for h2 < l2{
			res[i] = record2[h2]
			i++
			h2++
		}
	}else if h2 == l2{
		for h1 < l1{
			res[i] = record1[h1]
			i++
			h1++
		}
	}
	return res
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var record1 []int
	pre_visit1305(root1,&record1)
	var record2 []int
	pre_visit1305(root2,&record2)
	return merge_array(record1,record2)
}