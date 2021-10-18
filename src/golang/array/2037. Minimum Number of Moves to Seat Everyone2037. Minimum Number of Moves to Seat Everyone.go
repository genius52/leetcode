package array

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var steps int = 0
	var l int = len(seats)
	for i := 0;i < l;i++{
		steps += abs_int(seats[i] - students[i])
	}
	return steps
}