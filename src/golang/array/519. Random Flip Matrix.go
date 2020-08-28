package array

import (
	"math/rand"
	"time"
)

type Solution struct {
	rows int
	columns int
	zero_cnt int
	record map[int]int //save the numbers have been used before and it's index,the index number didn't be used.
}

func Constructor519(n_rows int, n_cols int) Solution {
	var obj Solution
	obj.rows = n_rows
	obj.columns = n_cols
	obj.zero_cnt = n_rows * n_cols
	obj.record = make(map[int]int)
	return obj
}

func (this *Solution) Flip() []int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(this.zero_cnt)
	this.zero_cnt--
	var res int
	pre_index,ok := this.record[r]
	if ok {
		res = pre_index
	}else{
		res = r
	}
	if pre,ok := this.record[this.zero_cnt];ok{
		this.record[r] = pre
	}else{
		this.record[r] = this.zero_cnt
	}
	return []int{res / this.columns,res % this.columns}
}

func (this *Solution) Reset()  {
	this.record = make(map[int]int)
	this.zero_cnt = this.rows * this.columns
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
