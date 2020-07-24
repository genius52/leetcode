package number

import (
	"strconv"
	"strings"
)

//Input: date = "2004-03-01"
//Output: 61
//date.length == 10
//date[4] == date[7] == '-', and all other date[i]'s are digits
//date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.
func DayOfYear(date string) int {
	var month_days map[string]int = make(map[string]int)
	month_days["01"] = 0
	month_days["02"] = month_days["01"] + 31
	month_days["03"] = month_days["02"] + 28
	month_days["04"] = month_days["03"] + 31
	month_days["05"] = month_days["04"] + 30
	month_days["06"] = month_days["05"] + 31
	month_days["07"] = month_days["06"] + 30
	month_days["08"] = month_days["07"] + 31
	month_days["09"] = month_days["08"] + 31
	month_days["10"] = month_days["09"] + 30
	month_days["11"] = month_days["10"] + 31
	month_days["12"] = month_days["11"] + 30
	record := strings.Split(date,"-")
	var days int = month_days[record[1]]
	if d,err := strconv.Atoi(record[2]);err == nil{
		days += d
	}
	if year,err := strconv.Atoi(record[0]);err == nil{
		if (year % 4 == 0 && year % 100 != 0){
			if(record[1] != "01" && record[1] != "02"){
				days++
			}
		}else if(year % 400 == 0){
			if(record[1] != "01" && record[1] != "02"){
				days++
			}
		}
	}
	return days
}