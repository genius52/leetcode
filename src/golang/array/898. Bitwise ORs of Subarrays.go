package array

//Input: [1,2,4]
//Output: 6
//Explanation:
//The possible results are 1, 2, 3, 4, 6, and 7.

//func combine_subarrayBitwiseORs(pos int,pre_val int,target_len int,A []int,record map[int]int){
//	if target_len == 0{
//		record[pre_val]++
//		return
//	}
//	if pos >= len(A){
//		return
//	}
//	//skip current
//	combine_subarrayBitwiseORs(pos + 1,pre_val,target_len,A,record)
//	//add current
//	combine_subarrayBitwiseORs(pos + 1,pre_val | A[pos],target_len - 1,A,record)
//}

//brute force
//func SubarrayBitwiseORs(A []int) int {
//	var l int = len(A)
//	var record map[int]int = make(map[int]int)
//	for size := 1;size <= l;size++{
//		for i := 0;i + size <= l;i++{
//			var total int = 0
//			for j := 0;j < size;j++{
//				total = total | A[i + j]
//			}
//			record[total]++
//		}
//	}
//	return len(record)
//}

func SubarrayBitwiseORs(A []int) int{
	var l int = len(A)
	var total_record map[int]bool = make(map[int]bool)
	var last_record map[int]bool = make(map[int]bool)
	for i := 0;i < l;i++{
		var cur_record map[int]bool = make(map[int]bool)
		cur_record[A[i]] = true
		for k,_ := range last_record{
			new_k := k | A[i]
			cur_record[new_k] = true
		}
		last_record = cur_record
		for k,_ := range last_record{
			total_record[k] = true
		}
	}
	return len(total_record)
}
