package number

import "sort"

//Input: people = [1,2], limit = 3
//Output: 1
//Explanation: 1 boat (1, 2)
//Input: people = [3,2,2,1], limit = 3
//Output: 3
//Explanation: 3 boats (1, 2), (2) and (3)
func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var l int = len(people)
	var res int = 0
	var begin int = 0
	for end := l - 1;end >= begin;end--{
		if people[end] + people[begin] <= limit{
			begin++
		}
		res++
	}
	return res
}