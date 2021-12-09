package tree

func find_path2096(root *TreeNode, val int, path *[]uint8) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	ret1 := find_path2096(root.Left, val, path)
	if ret1 {
		*path = append(*path, 'L')
		return true
	}
	ret2 := find_path2096(root.Right, val, path)
	if ret2 {
		*path = append(*path, 'R')
		return true
	}
	return false
}

func GetDirections(root *TreeNode, startValue int, destValue int) string {
	var start []uint8
	find_path2096(root, startValue, &start)
	var dest []uint8
	find_path2096(root, destValue, &dest)
	var l1 int = len(start)
	for i := 0; i < l1/2; i++ {
		start[i], start[l1-1-i] = start[l1-1-i], start[i]
	}
	var l2 int = len(dest)
	for i := 0; i < l2/2; i++ {
		dest[i], dest[l2-1-i] = dest[l2-1-i], dest[i]
	}
	var same_idx int = 0
	for same_idx < l1 && same_idx < l2 {
		if start[same_idx] == dest[same_idx] {
			same_idx++
		} else {
			break
		}
	}
	if l1-same_idx > 0 {
		start = start[same_idx:]
		for i := 0; i < l1-same_idx; i++ {
			start[i] = 'U'
		}
	} else {
		start = []uint8{}
	}
	dest = dest[same_idx:]
	res := string(start[:]) + string(dest[:])
	return res
}
