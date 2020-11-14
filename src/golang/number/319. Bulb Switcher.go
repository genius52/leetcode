package number

func BulbSwitch(n int) int {
	if(n == 1){
		return 1
	}
	var bulbs []bool = make([]bool,n)
	for interval := 1;interval <= n;interval++{
		for i := interval - 1;i < n;i += interval{
			bulbs[i] = !bulbs[i]
		}
	}
	var cnt int = 0
	for i := 0;i < n;i++{
		if(bulbs[i]){
			cnt++
		}
	}
	return cnt
}