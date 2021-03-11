package number

import "container/list"

//Input:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//Output:
//[9, 8, 6, 5, 3]

func MaxNumber(nums1 []int, nums2 []int, k int) []int{
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var l int = min_int(l1,l2)
	l = min_int(l,k)
	var max_num []int = make([]int,k)
	if l1 + l2 == k{
		return merge_array(nums1,nums2,k)

	}else{
		for i := 1;i <= min_int(k,l1);i++{
			j := k - i
			if l2 < j{
				continue
			}
			//j = min_int(j,l2)
			choose1 := max_array(nums1,i)
			choose2 := max_array(nums2,j)
			merge_data := merge_array(choose1,choose2,k)
			res := compareTwoArrays(merge_data,max_num)
			if res{
				copy(max_num,merge_data)
			}
		}
	}
	return max_num
}

func compareTwoArrays(nums1 []int,nums2 []int)bool {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	if l1 == 0{
		return true
	}
	if l2 == 0{
		return false
	}
	var l int = max_int(l1,l2)
	for i := 0;i < l;i++{
		if nums1[i] != nums2[i]{
			return nums1[i] > nums2[i]
		}
	}
	return true
}

func max_array(nums []int,k int)[]int{
	var l int = len(nums)
	var q list.List
	q.PushBack(nums[0])
	var visit int = 1
	for visit < l{
		for q.Len() > 0 && (q.Len() + l - visit) > k && visit < l && q.Back().Value.(int) < nums[visit]{
			q.Remove(q.Back())
		}
		if q.Len() < k{
			q.PushBack(nums[visit])
		}
		visit++
	}
	var res []int
	for i := 0;i < k;i++{
		res = append(res,q.Front().Value.(int))
		q.Remove(q.Front())
	}
	return res
}

func compare_two_array(nums1 []int,nums2 []int,idx1 *int,idx2* int)bool{
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	if l1 == 0{
		return false
	}
	if l2 == 0{
		return true
	}
	var max_len int = max_int(l1,l2)
	var k1 int = 0
	var k2 int = 0
	for i := 0;i < max_len;i++{
		if k1 >= l1{
			return false
		}
		if k2 >= l2{
			return true
		}
		if nums1[k1] != nums2[k2]{
			return nums1[k1] > nums2[k2]
		}
		k1++
		k2++
	}
	return true
}

func merge_array(nums1 []int,nums2 []int,k int)[]int{
	//var l1 int = len(nums1)
	//var l2 int = len(nums2)
	var res []int = make([]int,k)
	var index int = 0
	var i int = 0
	var j int = 0
	for index < k{
		if compare_two_array(nums1[i:],nums2[j:],&i,&j){
			res[index] = nums1[i]
			i++
		}else{
			res[index] = nums2[j]
			j++
		}
		index++
	}
	return res
}

//func dp_maxNumber(nums1 []int,nums2 []int,pos1 int,pos2 int,l1 int,l2 int,k int,cur_num int,memo [][][]int)int{
//	if k == 0 || (pos1 >= l1 && pos2 >= l2){
//		return cur_num
//	}
//	//if memo[k][pos1][pos2] != -1{
//	//	return memo[k][pos1][pos2]
//	//}
//	if pos1 >= l1{
//		choose := dp_maxNumber(nums1,nums2,pos1,pos2 + 1,l1,l2,k - 1,cur_num * 10 + nums2[pos2],memo)
//		var skip int = 0
//		if (l2 - pos2) > k{
//			skip = dp_maxNumber(nums1,nums2,pos1,pos2 + 1,l1,l2,k,cur_num,memo)
//		}
//		memo[k][pos1][pos2] = max_int_number(choose,skip)
//		return memo[k][pos1][pos2]
//	}
//	if pos2 >= l2{
//		choose := dp_maxNumber(nums1,nums2,pos1 + 1,pos2,l1,l2,k - 1,cur_num * 10 + nums1[pos1],memo)
//		var skip int = 0
//		if (l1 - pos1) > k{
//			skip = dp_maxNumber(nums1,nums2,pos1 + 1,pos2,l1,l2,k,cur_num,memo)
//		}
//		memo[k][pos1][pos2] = max_int_number(choose,skip)
//		return memo[k][pos1][pos2]
//	}
//	//choose num1 element
//	choose_1 := dp_maxNumber(nums1,nums2,pos1 + 1,pos2,l1,l2,k - 1,cur_num * 10 + nums1[pos1],memo)
//	//choose num2 element
//	choose_2 := dp_maxNumber(nums1,nums2,pos1,pos2 + 1,l1,l2,k - 1,cur_num * 10 + nums2[pos2],memo)
//	//skip num1
//	var skip1 int = 0
//	var skip2 int = 0
//	var rest_num_cnt = l1 - pos1 + l2 - pos2
//	if rest_num_cnt > k{
//		skip1 = dp_maxNumber(nums1,nums2,pos1 + 1,pos2,l1,l2,k,cur_num,memo)
//		skip2 = dp_maxNumber(nums1,nums2,pos1,pos2 + 1,l1,l2,k,cur_num,memo)
//	}
//	memo[k][pos1][pos2] = max_int_number(choose_1,choose_2,skip1,skip2)
//	return memo[k][pos1][pos2]
//}
//
//func MaxNumber(nums1 []int, nums2 []int, k int) []int {
//	var l1 int = len(nums1)
//	var l2 int = len(nums2)
//	var memo [][][]int = make([][][]int,k + 1)
//	for m := 0;m <= k;m++{
//		memo[m] = make([][]int,l1 + 1)
//		for i := 0;i <= l1;i++{
//			memo[m][i] = make([]int,l2 + 1)
//			for j := 0;j <= l2;j++{
//				memo[m][i][j] = -1
//			}
//		}
//	}
//
//	n := dp_maxNumber(nums1,nums2,0,0,l1,l2,k,0,memo)
//	var res []int
//	for n > 0{
//		mod := n % 10
//		n /= 10
//		res = append([]int{mod},res...)
//	}
//	return res
//}