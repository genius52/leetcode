package number

// 所有的钱都必须被分配。
// 每个儿童至少获得 1 美元。
// 没有人获得 4 美元。
// 返回 最多 有多少个儿童获得 恰好 8 美元
func DistMoney(money int, children int) int {
	if money < children {
		return -1
	}
	var eight_cnt int = children
	for eight_cnt > 0 {
		rest_children := children - eight_cnt
		rest_money := money - eight_cnt*8
		if (rest_money < 0) || (rest_children == 1 && rest_money == 4) {
			eight_cnt--
			continue
		}
		if (rest_money == 0 && rest_children == 0) || (rest_children != 0 && rest_money >= rest_children) {
			break
		}
		eight_cnt--
	}
	return eight_cnt
}
