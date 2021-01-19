package number

func CountPairs(deliciousness []int) int {
	var l int = len(deliciousness)
	var require []int = []int{1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536,131072,
		262144,524288,1048576,2097152}
	var res int = 0
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		for j := 0;j < 22;j++{
			need := require[j] - deliciousness[i]
			if cnt,ok := record[need];ok{
				res += cnt
				res %= 1000000007
			}
		}
		if _,ok := record[deliciousness[i]];ok{
			record[deliciousness[i]]++
		}else{
			record[deliciousness[i]] = 1
		}
	}
	return res
}