package number

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	var one_day_minutes int = 24 * 60
	var l int = len(timePoints)
	var data []int = make([]int,l)
	for i := 0;i < l;i++{
		time := strings.Split(timePoints[i],":")
		if hours,err := strconv.Atoi(time[0]);err == nil{
			if minutes,err2 := strconv.Atoi(time[1]);err2 == nil{
				data[i] = hours * 60 + minutes
			}
		}
	}
	sort.Ints(data)
	var min_minute int = math.MaxInt32
	for i := 1;i < l;i++{
		if data[i] - data[i - 1] < min_minute{
			min_minute = data[i] - data[i - 1]
		}
		if data[i - 1] + one_day_minutes - data[i] < min_minute{
			min_minute = data[i - 1] + one_day_minutes - data[i]
		}
	}
	if data[l - 1] - data[0] < min_minute{
		min_minute = data[l - 1] - data[0]
	}
	if data[0] + one_day_minutes - data[l - 1] < min_minute{
		min_minute = data[0] + one_day_minutes - data[l - 1]
	}
	return min_minute
}