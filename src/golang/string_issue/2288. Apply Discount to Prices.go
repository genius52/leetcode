package string_issue

import (
	"fmt"
	"strconv"
	"strings"
)

func DiscountPrices(sentence string, discount int) string {
	var data []string = strings.Split(sentence," ")
	var ss strings.Builder
	for i := 0;i < len(data);i++{
		if i != 0{
			ss.WriteString(" ")
		}
		if data[i][0] == '$'{
			n, err := strconv.ParseFloat(data[i][1:], 64)
			if err == nil{
				n2 := n * float64(100 - discount) / float64(100.0)
				f,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", n2), 64)
				val := strconv.FormatFloat(f, 'f', 2, 64)
				ss.WriteString("$")
				ss.WriteString(val)
			}else{
				ss.WriteString(data[i])
			}
		}else{
			ss.WriteString(data[i])
		}
	}
	return ss.String()
}