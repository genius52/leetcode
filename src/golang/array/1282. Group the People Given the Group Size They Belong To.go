package array

//Input: groupSizes = [3,3,3,3,3,1,3]
//Output: [[5],[0,1,2],[3,4,6]]
//Explanation:
//Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
func groupThePeople2(groupSizes []int) [][]int{
	var l int = len(groupSizes)
	var size_index map[int][]int = make(map[int][]int)
	var res [][]int
	for i := 0;i < l;i++{
		size := groupSizes[i]
		if _,ok := size_index[size];!ok{
			size_index[size] = []int{i}
		}else{
			size_index[size] = append(size_index[size],i)
		}
		if len(size_index[size]) == size{
			res = append(res,size_index[size])
			delete(size_index,size)
		}
	}
	return res
}

func groupThePeople(groupSizes []int) [][]int {
	//save the array index by size
	var record map[int][] int = make(map[int][]int)
	for i,s := range groupSizes{
		if _,ok := record[s];ok{
			record[s] = append(record[s], i)
		}else{
			record[s] = []int{i}
		}
	}
	var res [][]int
	for length,nums := range record{
		var collection []int
		for i := 0;i < len(nums);i++{
			collection = append(collection, nums[i])
			if i % length == length - 1{
				res = append(res, collection)
				collection = []int{}
			}
		}
	}
	return res
}