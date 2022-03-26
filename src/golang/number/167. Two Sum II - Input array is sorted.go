package number

func twoSum167(numbers []int, target int) []int {
	var l int = len(numbers)
	var res []int
	var left int = 0
	var right int = l - 1
	for left < right{
		if numbers[left] + numbers[right] < target{
			left++
		}else if numbers[left] + numbers[right] > target{
			right--
		}else{
			res = []int{left+1,right+1}
			break
		}
	}
	return res
}

//func twoSum167(numbers []int, target int) []int {
//	var l int = len(numbers)
//	var record map[int]int = make(map[int]int)
//	var res []int
//	for i := 0;i < l;i++{
//		need := target - numbers[i]
//		if pos,ok := record[need];ok{
//			res = append(res,pos)
//			res = append(res,i + 1)
//			return res
//		}
//		record[numbers[i]] = i + 1
//	}
//	return res
//}