package number


//933
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
type RecentCounter struct {
	record []int
}

func Constructor933() RecentCounter {
	var obj RecentCounter
	return obj
}

func (this *RecentCounter) Ping(t int) int {
	this.record = append(this.record,t)
	low := 0
	l := len(this.record)
	high := l - 1
	target := t - 3000
	mid := 0
	//find the postion of first num large or equal the target
	for low <= high{
		mid = (low + high)/2
		if this.record[mid] == target{
			break
		}
		if this.record[mid] < target{
			low = mid + 1
			if mid < (l - 1) && this.record[mid + 1] >= target{
				return len(this.record) - mid - 1
			}
		}else{
			high = mid - 1
			if mid > 0 && this.record[mid - 1] < target{
				return len(this.record) - mid
			}
		}
	}
	return len(this.record) - mid
}