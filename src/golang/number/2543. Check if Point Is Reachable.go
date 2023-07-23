package number

//给你一个无穷大的网格图。一开始你在(1, 1)，你需要通过有限步移动到达点(targetX, targetY)。
//
//每一步，你可以从点(x, y)移动到以下点之一：
//
//(x, y - x)
//(x - y, y)
//(2 * x, y)
//(x, 2 * y)

//(x, y + x)
//(x + y, y)
//(x/2, y) 如果x是偶数
//(x, y/2) 如果y是偶数

func isReachable(targetX int, targetY int) bool {
	r := gcd(targetX, targetY)
	return r&(r-1) == 0
}
