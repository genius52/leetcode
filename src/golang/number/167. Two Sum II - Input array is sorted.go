package number

func twoSum167(numbers []int, target int) []int {
	var l int = len(numbers)
	var record map[int]int = make(map[int]int)
	var res []int
	for i := 0;i < l;i++{
		need := target - numbers[i]
		if pos,ok := record[need];ok{
			res = append(res,pos)
			res = append(res,i + 1)
			return res
		}
		record[numbers[i]] = i + 1
	}
	return res
}