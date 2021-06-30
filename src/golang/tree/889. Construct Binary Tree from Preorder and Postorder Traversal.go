package tree

//Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//Output: [1,2,3,4,5,6,7]

func build_constructFromPrePost(pre []int,pre_start int,pre_end int,post []int,post_start int,post_end int)*TreeNode{
	if pre_end >= len(pre) || post_end >= len(post) || pre_start > pre_end || post_start > post_end{
		return nil
	}
	var node TreeNode
	node.Val = pre[pre_start]
	if pre_start == pre_end && post_start == post_end{
		return &node
	}
	var left_end_in_post int = post_end
	for post[left_end_in_post] != pre[pre_start + 1] && left_end_in_post > post_start{
		left_end_in_post--
	}
	left_end_in_pre := pre_start + left_end_in_post - post_start + 1
	node.Left = build_constructFromPrePost(pre,pre_start + 1, left_end_in_pre,post,post_start,left_end_in_post)
	node.Right = build_constructFromPrePost(pre,left_end_in_pre + 1,pre_end,post,left_end_in_post + 1,post_end - 1)
	return &node
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	l := len(pre)
	if l == 0{
		return nil
	}
	return build_constructFromPrePost(pre,0,l - 1,post,0,l - 1)
}


//889
//var pre_cnt int = 0
//var post_cnt int = 0
//func constructFromPrePost(pre []int, post []int) *TreeNode {
//	var node *TreeNode = new(TreeNode)
//	node.Val = pre[pre_cnt]
//	pre_cnt++
//	if node.Val != post[post_cnt] {
//		node.Left = constructFromPrePost(pre,post)
//	}
//	if node.Val != post[post_cnt] {
//		node.Right = constructFromPrePost(pre,post)
//	}
//	post_cnt++
//	return node
//}