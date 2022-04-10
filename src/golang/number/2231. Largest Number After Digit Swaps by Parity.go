package number

func largestInteger(num int) int {
	var data []int
	for num > 0{
		data = append([]int{num%10},data...)
		num /= 10
	}
	for i := 0;i < len(data);i++{
		for j := i + 1;j < len(data);j++{
			if (data[i] | 1 == data[i] && data[j] | 1 == data[j]) || (data[i] | 1 != data[i] && data[j] | 1 != data[j]){
				if data[i] < data[j]{
					data[i],data[j] = data[j],data[i]
				}
			}
		}
	}
	var res int = 0
	for i := 0;i < len(data);i++{
		res *= 10
		res += data[i]
	}
	return res
}