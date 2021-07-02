package array

import "container/list"

type Price_duration struct {
	price int
	duration int
}

type StockSpanner struct {
	q list.List
}

func Constructor901() StockSpanner {
	var obj StockSpanner
	return obj
}

func (this *StockSpanner) Next(price int) int {
	var res int = 1
	for this.q.Len() > 0{
		var cur Price_duration = this.q.Back().Value.(Price_duration)
		if cur.price <= price{
			res += cur.duration
			this.q.Remove(this.q.Back())
		}else{
			break
		}
	}
	var obj Price_duration
	obj.price = price
	obj.duration = res
	this.q.PushBack(obj)
	return res
}

//type StockSpanner struct {
//	prices []int
//	dp []int
//	length int
//}
//
//func Constructor901() StockSpanner {
//	var obj StockSpanner
//	return obj
//}
//
//
//func (this *StockSpanner) Next(price int) int {
//	if this.length > 0 && price >= this.prices[this.length - 1]{
//		visit := this.length - 1
//		for visit >=0 {
//			if price >= this.prices[visit]{
//				visit = visit - this.dp[visit]
//			}else{
//				break
//			}
//		}
//		this.dp = append(this.dp,this.length - visit)
//	}else{
//		this.dp = append(this.dp,1)
//	}
//	this.prices = append(this.prices,price)
//	this.length++
//	return this.dp[this.length - 1]
//}