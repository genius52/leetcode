package diagram

func CheckStraightLine(coordinates [][]int) bool {
	var l int = len(coordinates)
	if l <= 2{
		return true
	}
	for i := 1;i < l - 1;i++{
		val1 := (coordinates[i][1] - coordinates[i - 1][1]) * (coordinates[i][0] - coordinates[i + 1][0])
		val2 := (coordinates[i][1] - coordinates[i + 1][1]) * (coordinates[i][0] - coordinates[i - 1][0])
		if  val1 != val2 {
			return false
		}
	}
	return true
}