package array

type StockSpanner struct {
	prices []int
	dp []int
	length int
}

func Constructor901() StockSpanner {
	var obj StockSpanner
	return obj
}


func (this *StockSpanner) Next(price int) int {
	if this.length > 0 && price >= this.prices[this.length - 1]{
		//this.dp = append(this.dp,this.dp[this.length - 1] + 1)
		visit := this.length - 1
		for visit >=0 {
			if price >= this.prices[visit]{
				visit = visit - this.dp[visit]
			}else{
				break
			}
		}
		this.dp = append(this.dp,this.length - visit)
	}else{
		this.dp = append(this.dp,1)
	}
	this.prices = append(this.prices,price)
	this.length++
	return this.dp[this.length - 1]
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */