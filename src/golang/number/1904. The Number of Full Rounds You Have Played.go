package number

import (
	"strconv"
	"strings"
)

func NumberOfRounds(startTime string, finishTime string) int {
	if startTime == finishTime{
		return 0
	}
	var s1 []string = strings.Split(startTime,":")
	var s2 []string = strings.Split(finishTime,":")
	starthour,_ := strconv.Atoi(s1[0])
	startminute,_ := strconv.Atoi(s1[1])

	finishhour,_ := strconv.Atoi(s2[0])
	finishminute,_ := strconv.Atoi(s2[1])

	var res int = 0
	if (starthour > finishhour) || (starthour == finishhour && startminute > finishminute){
		if startminute > 0 && startminute < 15{
			startminute = 15
		}else if startminute > 15 && startminute < 30{
			startminute = 30
		}else if startminute > 30 && startminute < 45{
			startminute = 45
		}else if startminute > 45 && startminute <= 59{
			starthour = starthour + 1
			startminute = 0
		}
		if finishminute > 0 && finishminute < 15{
			finishminute = 0
		}else if finishminute > 15 && finishminute < 30{
			finishminute = 15
		}else if finishminute > 30 && finishminute < 45{
			finishminute = 30
		}else if finishminute > 45 && finishminute <= 59{
			finishminute = 45
		}
		res += (finishhour + 24 - starthour) * 4
		res += (finishminute - startminute) / 15
	}else{
		if startminute > 0 && startminute < 15{
			startminute = 15
		}else if startminute > 15 && startminute < 30{
			startminute = 30
		}else if startminute > 30 && startminute < 45{
			startminute = 45
		}else if startminute > 45 && startminute <= 59{
			starthour = starthour + 1
			startminute = 0
		}
		if finishminute > 0 && finishminute < 15{
			finishminute = 0
		}else if finishminute > 15 && finishminute < 30{
			finishminute = 15
		}else if finishminute > 30 && finishminute < 45{
			finishminute = 30
		}else if finishminute > 45 && finishminute <= 59{
			finishminute = 45
		}
		res += (finishhour - starthour) * 4
		res += (finishminute - startminute) / 15
	}
	if res < 0{
		return 0
	}
	return res
}