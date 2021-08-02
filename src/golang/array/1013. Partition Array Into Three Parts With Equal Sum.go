package array

func canThreePartsEqualSum(A []int) bool {
	var total int = 0
	var l int = len(A)
	for i:= 0;i < l;i++{
		total += A[i]
	}
	if total % 3 != 0{
		return false
	}
	var target int = total/3
	var sum int = 0
	var find_cnt = 0
	for i:= 0;i < l;i++{
		sum += A[i]
		if sum == target{
			find_cnt++
			if find_cnt == 2 {
				if i != (l - 1){
					return true
				}else{
					return false
				}
			}
			sum = 0
		}
	}
	return false
}