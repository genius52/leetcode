package number

import "sort"

//Input: a = 4, b = 4, c = 6
//Output: 7
//Explanation: The starting state is (4, 4, 6). One optimal set of moves is:
//- Take from 1st and 2nd piles, state is now (3, 3, 6)
//- Take from 1st and 3rd piles, state is now (2, 3, 5)
//- Take from 1st and 3rd piles, state is now (1, 3, 4)
//- Take from 1st and 3rd piles, state is now (0, 3, 3)
//- Take from 2nd and 3rd piles, state is now (0, 2, 2)
//- Take from 2nd and 3rd piles, state is now (0, 1, 1)
//- Take from 2nd and 3rd piles, state is now (0, 0, 0)
func mid_int_number(a,b,c int)int{
	var data []int = []int{a,b,c}
	sort.Ints(data)
	return data[1]
}
//4,5,6
//10,11,12
//5,6,12    5
//5,0,6     6
//0,0,1     5
func maximumScore(a int, b int, c int) int {
	big := max_int_number(a,b,c)
	small := min_int_number(a,b,c)
	middle := mid_int_number(a,b,c)
	if big - middle <= small{
		return (small + middle + big)/2
	}
	return small + middle
}