package number

import "strconv"

func haveConflict(event1 []string, event2 []string) bool {
	start1_hour,_ := strconv.Atoi(event1[0][:2])
	start1_minute,_ := strconv.Atoi(event1[0][3:])
	start2_hour,_ := strconv.Atoi(event2[0][:2])
	start2_minute,_ := strconv.Atoi(event2[0][3:])
	end1_hour,_ := strconv.Atoi(event1[1][:2])
	end1_minute,_ := strconv.Atoi(event1[1][3:])
	end2_hour,_ := strconv.Atoi(event2[1][:2])
	end2_minute,_ := strconv.Atoi(event2[1][3:])
	start1_time := start1_hour * 60 + start1_minute
	end1_time := end1_hour * 60 + end1_minute
	start2_time := start2_hour * 60 + start2_minute
	end2_time := end2_hour * 60 + end2_minute
	var start int = start1_time
	if start < start2_time{
		start = start2_time
	}
	var end int = end1_time
	if end > end2_time{
		end = end2_time
	}
	return start <= end
}