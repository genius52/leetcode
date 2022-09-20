package number

import "strconv"

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	start_month1,_ := strconv.Atoi(arriveAlice[:2])
	start_day1,_ := strconv.Atoi(arriveAlice[3:])
	leave_month1,_ := strconv.Atoi(leaveAlice[:2])
	leave_days1,_ := strconv.Atoi(leaveAlice[3:])
	Alice_start_days := start_month1 * 100 + start_day1
	Alice_leave_days := leave_month1 * 100 + leave_days1

	start_month2,_ := strconv.Atoi(arriveBob[:2])
	start_day2,_ := strconv.Atoi(arriveBob[3:])
	leave_month2,_ := strconv.Atoi(leaveBob[:2])
	leave_days2,_ := strconv.Atoi(leaveBob[3:])
	Bob_start_days := start_month2 * 100 + start_day2
	Bob_leave_days := leave_month2 * 100 + leave_days2

	if start_month1 > leave_month2 || start_month2 > leave_month1{
		return 0
	}
	var month_days [12]int = [12]int{31,28,31,30,31,30,31,31,30,31,30,31}
	var res int = 0
	for i := max_int(start_month1,start_month2);i <= min_int(leave_month1,leave_month2);i++{
		if i == max_int(start_month1,start_month2) || i == min_int(leave_month1,leave_month2){
			for j := 1;j <= month_days[i - 1];j++{
				days := i * 100 + j
				if days >= Alice_start_days && days <= Alice_leave_days && days >= Bob_start_days && days <= Bob_leave_days{
					res++
				}
			}
		}else{
			res += month_days[i - 1]
		}
	}
	return res
}