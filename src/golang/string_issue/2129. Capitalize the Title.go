package string_issue

import "strings"

func capitalizeTitle(title string) string {
	var data []string = strings.Split(title, " ")
	var res strings.Builder
	for _, d := range data {
		d = strings.ToLower(d)
		if res.Len() > 0 {
			res.WriteByte(' ')
		}
		if len(d) <= 2 {
			res.WriteString(d)
		} else {
			res.WriteByte(d[0] - 32)
			res.WriteString(d[1:])
		}
	}
	return res.String()
}

//func capitalizeTitle(title string) string {
//	var l int = len(title)
//	var res strings.Builder
//	var pre uint8 = ' '
//	for i := 0; i < l; i++ {
//		c := title[i]
//		if pre == ' ' {
//			if c >= 'a' && title[i] <= 'z' {
//				c -= 32
//			}
//		} else {
//			if c >= 'A' && title[i] <= 'Z' {
//				c += 32
//			}
//		}
//		res.WriteByte(c)
//		pre = title[i]
//	}
//	return res.String()
//}
