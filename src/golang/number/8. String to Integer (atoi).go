package number

import (
	"math"
	"strconv"
	"strings"
)

//Input: "4193 with words"
//Output: 4193
func MyAtoi(str string) int {
	str = strings.Trim(str," ")
	var l int = len(str)
	if(l == 0){
		return 0
	}
	if(str[0] != '-' && str[0] != '+' && (str[0] > '9' || str[0] < '0')){
		return 0
	}
	var pos int = 0
	var negative bool = false
	if(str[0] == '-'){
		negative = true
		pos++
	}else if(str[0] == '+'){
		pos++
	}

	for(pos < l){
		if(str[pos] >= '0' && str[pos] <= '9'){
			pos++
			continue
		}else{
			break
		}
	}
	if n,err := strconv.Atoi(str[:pos]);err == nil{
		if(n < math.MinInt32){
			return math.MinInt32
		}
		if(n > math.MaxInt32){
			return math.MaxInt32
		}
		return n
	}else{
		if(strings.Contains(err.Error(), "value out of range")){
			if(negative){
				return math.MinInt32
			}else{
				return math.MaxInt32
			}
		}
		return 0
	}
}