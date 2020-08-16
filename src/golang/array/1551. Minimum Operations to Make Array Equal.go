package array

func MinOperations(n int) int {
	var steps int = 0
	var biggest int = (n - 1) * 2 + 1
	var standard int = (biggest + 1)/2
	for i := 1;i <= standard;i += 2{
		steps = steps + (standard - i)
	}
	return steps
}
