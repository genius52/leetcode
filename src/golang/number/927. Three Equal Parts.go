package number

//Input: arr = [1,0,1,0,1]
//Output: [0,3]

//Input: arr = [1,1,0,1,1]
//Output: [-1,-1]

//Input: arr = [1,1,0,0,1]
//Output: [0,2]
func ThreeEqualParts(arr []int) []int {
	var l int = len(arr)
	if l < 3{
		return []int{-1,-1}
	}
	var start int = 0
	for start < l{
		if arr[start] == 0{
			start++
		}else{
			break
		}
	}
	if start == l{
		return []int{0,2}
	}
	var one_cnt int = 0
	for i := start;i < l;i++{
		if arr[i] == 1{
			one_cnt++
		}
	}
	if one_cnt % 3 != 0{
		return []int{-1,-1}
	}
	var target_one_ont int = one_cnt / 3
	var left int = start
	var cur_one_cnt int = 0
	for cur_one_cnt < target_one_ont{
		if arr[left] == 1{
			cur_one_cnt++
		}
		left++
	}
	var first int = left - 1 //在1的位置
	var postfix_len1 int = 0
	for arr[left] == 0{
		left++
		postfix_len1++
	}
	var idx1 int = start
	var idx2 int = left
	cur_one_cnt = 0
	for cur_one_cnt < target_one_ont{
		if arr[idx1] != arr[idx2]{
			return []int{-1,-1}
		}
		if arr[idx1] == 1{
			cur_one_cnt++
		}
		idx1++
		idx2++
	}
	var second int = idx2 //在1后面的位置
	var postfix_len2 int = 0
	var idx3 int = idx2
	for arr[idx3] == 0{
		idx3++
		postfix_len2++
	}
	idx1 = start
	cur_one_cnt = 0
	for cur_one_cnt < target_one_ont{
		if arr[idx1] != arr[idx3]{
			return []int{-1,-1}
		}
		if arr[idx1] == 1{
			cur_one_cnt++
		}
		idx1++
		idx3++
	}
	postfix_len3 := l - idx3
	//12331 xxx   1212 xxx 13131 xxx
	if postfix_len1 < postfix_len3 || postfix_len2 < postfix_len3{
		return []int{-1,-1}
	}
	if postfix_len3 > 0{
		first += postfix_len3
		second += postfix_len3
	}
	return []int{first,second}
}