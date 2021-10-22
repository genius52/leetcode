package array

import "container/list"

//stack
func finalPrices(prices []int) []int {
	var l int = len(prices)
	var res []int = make([]int,l)
	var q list.List
	for i := l - 1;i >= 0;i--{
		for q.Len() > 0 && q.Back().Value.(int) > prices[i]{
			q.Remove(q.Back())
		}
		if q.Len() > 0{
			res[i] = prices[i] - q.Back().Value.(int)
		}else{
			res[i] = prices[i]
		}
		q.PushBack(prices[i])
	}
	return res
}

//brute force
func FinalPrices(prices []int) []int {
	var l int = len(prices)
	var res []int = make([]int,l)
	res[l - 1] = prices[l - 1]
	for i := l - 2;i >= 0;i--{
		var idx int = -1
		for j := i + 1;j < l;j++{
			if prices[j] <= prices[i]{
				idx = j
				break
			}
		}
		if idx == -1{
			res[i] = prices[i]
		}else{
			res[i] = prices[i] - prices[idx]
		}
	}
	return res
}