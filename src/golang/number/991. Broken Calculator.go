package number

import "math"

func dp_brokenCalc(X int,Y int,memo map[int]int,visited map[int]bool)int{
	if X <= 0{
		return math.MaxInt32
	}
	if cnt,ok := memo[X];ok{
		return cnt
	}
	if _,ok := visited[X];ok{
		return math.MaxInt32
	}
	visited[X] = true
	var res int
	if X == Y{
		res = 0
	}else if X > Y{
		res = 1 + dp_brokenCalc(X - 1,Y,memo,visited)
	}else if X < Y{
		decrease := 1 + dp_brokenCalc(X - 1,Y,memo,visited)
		multiple := 1 + dp_brokenCalc(X * 2,Y,memo,visited)
		res = int(math.Min(float64(multiple),float64(decrease)))
	}
	memo[X] = res
	return res
}

func BrokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}
	var steps int = 0
	var i int = X
	var j int = Y
	for i < j{
		i *= 2
		steps++
	}
	if i != j{
		steps = steps + i - j
	}
	return 1
}