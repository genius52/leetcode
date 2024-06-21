package number

func countCompleteDayPairs2(hours []int) int64 {
	var res int64 = 0
	var cnt map[int]int = make(map[int]int)
	for i := 0; i < len(hours); i++ {
		need := (24 - hours[i]%24) % 24
		if cnt2, ok := cnt[need]; ok {
			res += int64(cnt2)
		}
		cnt[hours[i]%24]++
	}
	return res
}
