package array

//Sn=na1+n(n-1)d/2
func MinOperations(n int) int{
	cnt := n / 2
	var first int = 2
	if n | 1 != n{
		first = 1
	}
	return (first * cnt + cnt * (cnt - 1) * 2/2)
}

//func MinOperations(n int) int {
//	var steps int = 0
//	var biggest int = (n - 1) * 2 + 1
//	var standard int = (biggest + 1)/2
//	for i := 1;i <= standard;i += 2{
//		steps = steps + (standard - i)
//	}
//	return steps
//}
