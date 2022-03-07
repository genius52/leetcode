package string_issue

import "strconv"

func CellsInRange(s string) []string {
	var start_c uint8 = s[0]
	var start_r uint8 = s[1] - '0'
	var end_c uint8 = s[3]
	var end_r uint8 = s[4] - '0'
	var res []string
	for i := 0;i <= int(end_c - start_c);i++{
		col := string(start_c + uint8(i))
		for j := 0;j <= int(end_r - start_r);j++{
			cur := col + strconv.Itoa(int(start_r) + j)
			res = append(res,cur)
		}
	}
	return res
}