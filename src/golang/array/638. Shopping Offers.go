package array

//Input: [2,3,4], [[1,1,0,4],[2,2,1,9]], [1,2,1]
//Output: 11
//Explanation:
//The price of A is $2, and $3 for B, $4 for C.
//You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C.
//You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer #1), and $3 for 1B, $4 for 1C.
//You cannot add more items, though only $9 for 2A ,2B and 1C.
func dp_shoppingOffers(price []int,special [][]int,needs []int)int{
	var no_need bool = true
	for _,n := range needs{
		if n != 0 {
			no_need = false
			break
		}
	}
	if no_need{
		return 0
	}
	var simple_cost int = 0
	for i,p := range price{
		simple_cost += needs[i] * p
	}
	var min_cost int = simple_cost
	var l int = len(price)
	for _,s := range special{
		external := false
		var rest_needs []int = make([]int,l)
		copy(rest_needs,needs)
		for i := 0;i < l;i++{
			rest_needs[i] -= s[i]
			if rest_needs[i] < 0{
				external = true
				break
			}
		}
		if 	external {
			continue
		}
		cost := s[l] + dp_shoppingOffers(price,special,rest_needs)
		if cost < min_cost{
			min_cost = cost
		}
	}
	return min_cost
}

func ShoppingOffers(price []int, special [][]int, needs []int) int {
	return dp_shoppingOffers(price,special,needs)
}