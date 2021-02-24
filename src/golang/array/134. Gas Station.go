package array

// [1,2,3,4,5]
// [3,4,5,1,2]

//	{2,3,4}
//	{3,4,3}
func canCompleteCircuit(gas []int, cost []int) int {
	var start_point int = -1
	var l = len(gas)
	for begin := 0; begin <= len(gas);begin++ {
		trace := begin + 1
		cur_gas := 0
		for begin !=  trace % l {
			cur_gas += gas[(trace-1) % l]
			if cur_gas < cost[(trace-1) % l]{
				break
			}else{
				cur_gas -= cost[(trace-1) % l]
			}
			trace++
		}
		cur_gas += gas[(trace-1) % l]
		if begin == trace % l && cur_gas >= cost[(trace-1) % l]{
			return begin
		}
	}
	return start_point
}