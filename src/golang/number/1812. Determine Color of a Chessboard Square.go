package number

func squareIsWhite(coordinates string) bool {
	c := coordinates[0] - 'a'
	n := coordinates[1] - '1'
	return !(c&1 == n&1)
}

//func squareIsWhite(coordinates string) bool {
//	c := coordinates[0] - 'a'
//	n := coordinates[1] - '1'
//	if c|1 == c {
//		if n|1 == n {
//			return false
//		} else {
//			return true
//		}
//	} else {
//		if n|1 == n {
//			return true
//		} else {
//			return false
//		}
//	}
//}
