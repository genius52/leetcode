package number

//ATM()初始化 ATM 对象。
//void deposit(int[] banknotesCount)分别存入$20，$50，$100，$200和$500钞票的数目。
//int[] withdraw(int amount)返回一个长度为5的数组，分别表示$20，$50，$100，$200和$500钞票的数目，
//并且更新 ATM 机里取款后钞票的剩余数量。如果无法取出指定数额的钱，请返回[-1]（这种情况下 不取出任何钞票）

type ATM struct {
	cnt20 int
	cnt50 int
	cnt100 int
	cnt200 int
	cnt500 int
}


func Constructor2242() ATM {
	var obj ATM
	return obj
}

func (this *ATM) Deposit(banknotesCount []int)  {
	this.cnt20 += banknotesCount[0]
	this.cnt50 += banknotesCount[1]
	this.cnt100 += banknotesCount[2]
	this.cnt200 += banknotesCount[3]
	this.cnt500 += banknotesCount[4]
}


func (this *ATM) Withdraw(amount int) []int {
	var n int = amount
	if this.cnt500 > 0 {
		need := n / 500
		n -= min_int(need, this.cnt500) * 500
	}
	if this.cnt200 > 0{
		need := n / 200
		n -= min_int(need, this.cnt200) * 200
	}
	if this.cnt100 > 0{
		need := n / 100
		n -= min_int(need, this.cnt100) * 100
	}
	if this.cnt50 > 0{
		need := n / 50
		n -= min_int(need, this.cnt50) * 50
	}
	if this.cnt20 > 0{
		need := n / 20
		n -= min_int(need, this.cnt20) * 20
	}
	if n != 0{
		return []int{-1}
	}
	var res []int = make([]int,5)
	if this.cnt500 > 0 {
		need := amount / 500
		amount -= min_int(need, this.cnt500) * 500
		res[4] = min_int(need, this.cnt500)
		this.cnt500 -= res[4]
	}
	if this.cnt200 > 0{
		need := amount / 200
		amount -= min_int(need, this.cnt200) * 200
		res[3] = min_int(need, this.cnt200)
		this.cnt200 -= res[3]
	}
	if this.cnt100 > 0{
		need := amount / 100
		amount -= min_int(need, this.cnt100) * 100
		res[2] = min_int(need, this.cnt100)
		this.cnt100 -= res[2]
	}
	if this.cnt50 > 0{
		need := amount / 50
		amount -= min_int(need, this.cnt50) * 50
		res[1] = min_int(need, this.cnt50)
		this.cnt50 -= res[1]
	}
	if this.cnt20 > 0{
		need := amount / 20
		amount -= min_int(need, this.cnt20) * 20
		res[0] = min_int(need, this.cnt20)
		this.cnt20 -= res[0]
	}
	return res
}