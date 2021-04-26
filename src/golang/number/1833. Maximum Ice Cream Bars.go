package number

import "sort"

//Input: costs = [1,3,2,4,1], coins = 7
//Output: 4
//Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var res int = 0
	for _,c := range costs{
		if coins >= c{
			coins -= c
			res++
		}else{
			break
		}
	}
	return res
}