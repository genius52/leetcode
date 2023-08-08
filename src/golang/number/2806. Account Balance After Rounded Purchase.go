package number

func accountBalanceAfterPurchase(purchaseAmount int) int {
	m := purchaseAmount % 10
	if m < 5 {
		return 100 - purchaseAmount/10*10
	} else {
		return 100 - (purchaseAmount/10+1)*10
	}
}
