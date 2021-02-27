package number

//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28
//    ...
func titleToNumber(s string) int {
	var total int = 0
	for _,c := range s{
		total *= 26
		total += int(c - 'A' + 1)
	}
	return total
}