package number

type FrequencyTracker struct {
	num_cnt map[int]int
	cnt_num map[int]map[int]bool
}

func Constructor2671() FrequencyTracker {
	var obj FrequencyTracker
	obj.num_cnt = make(map[int]int)
	obj.cnt_num = make(map[int]map[int]bool)
	return obj
}

func (this *FrequencyTracker) Add(number int) {
	if cnt, ok := this.num_cnt[number]; ok {
		this.num_cnt[number]++
		delete(this.cnt_num[cnt], number)
		if len(this.cnt_num[cnt]) == 0 {
			this.cnt_num[cnt] = nil
		}
		if this.cnt_num[cnt+1] == nil {
			this.cnt_num[cnt+1] = make(map[int]bool)
		}
		this.cnt_num[cnt+1][number] = true
	} else {
		this.num_cnt[number] = 1
		if this.cnt_num[1] == nil {
			this.cnt_num[1] = make(map[int]bool)
		}
		this.cnt_num[1][number] = true
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if cnt, ok := this.num_cnt[number]; ok {
		this.num_cnt[number]--
		if this.num_cnt[number] == 0 {
			delete(this.num_cnt, number)
		}
		delete(this.cnt_num[cnt], number)
		if len(this.cnt_num[cnt]) == 0 {
			this.cnt_num[cnt] = nil
		}
		if cnt != 1 {
			if this.cnt_num[cnt-1] == nil {
				this.cnt_num[cnt-1] = make(map[int]bool)
			}
			this.cnt_num[cnt-1][number] = true
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	if _, ok := this.cnt_num[frequency]; ok {
		if len(this.cnt_num[frequency]) > 0 {
			return true
		}
	}
	return false
}
