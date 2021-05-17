package number

//Return an array containing [crashTime, memory1crash, memory2crash],
//where crashTime is the time (in seconds) when the program crashed and memory1crash and memory2crash
//are the available bits of memory in the first and second sticks respectively.

//Input: memory1 = 2, memory2 = 2
//Output: [3,1,0]
func MemLeak(memory1 int, memory2 int) []int {
	var res []int = make([]int,3)
	var use_memo int = 1
	var time int = 1
	for memory1 >= use_memo || memory2 >= use_memo{
		if memory1 >= memory2{
			memory1 -= use_memo
		}else{
			memory2 -= use_memo
		}
		use_memo++
		time++
	}
	res[0] = time
	res[1] = memory1
	res[2] = memory2
	return res
}