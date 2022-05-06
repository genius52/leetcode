package number

import (
	"math/rand"
	"sort"
)


type solution struct {
	record [][2]int//Use sum []int instead
	sum int
}


func Constructor528(w []int) solution {
	var obj solution
	for idx,val := range w{
		obj.sum += val
		obj.record = append(obj.record,[2]int{obj.sum,idx})
	}
	return obj
}


func (this *solution) PickIndex() int {
	rand_val := rand.Intn(this.sum + 1) % (this.sum + 1)
	find_idx := sort.Search(len(this.record), func(i int) bool {
		return this.record[i][0] > rand_val
	})
	if find_idx != len(this.record){
		return this.record[find_idx][1]
	}else{
		return -1
	}
}
