package number

func countBalls(lowLimit int, highLimit int) int {
	var record map[int]int = make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		n := i
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		record[sum]++
	}
	var max_balls int = 0
	for _, v := range record {
		if v > max_balls {
			max_balls = v
		}
	}
	return max_balls
}
