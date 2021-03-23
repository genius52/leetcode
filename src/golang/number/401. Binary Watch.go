package number

import (
	"fmt"
	"math"
	"strconv"
)

func int_to_time(time []int) (bool,string){
	hour := 0
	minute := 0
	for i := 0 ;i <10;i++{
		if time[i] == 1 {
			if i >=6 {
				hour += time[i] * int(math.Pow(float64(2),float64(i-6)))
				if hour >= 12{
					return false,""
				}
			}else{
				minute += time[i] * int(math.Pow(float64(2),float64(i)))
				if minute >= 60{
					return false,""
				}
			}
		}
	}
	var res string = fmt.Sprintf("%d:%02d" ,hour,minute)
	return true,res
}

func combine(select_data []int,res *[]string,selected_num int,target_num int,cur_pos int,end_pos int){
	if nil == select_data {
		select_data = make([]int,10,10)
	}
	if selected_num >= target_num {
		if ok,val := int_to_time(select_data);ok{
			*res = append(*res,val)
		}
		return
	}
	if cur_pos >= end_pos{
		return
	}
	select_data[cur_pos] = 1
	combine(select_data,res,selected_num + 1,target_num,cur_pos+1,end_pos)
	select_data[cur_pos] = 0
	combine(select_data,res,selected_num,target_num,cur_pos+1,end_pos)
}

func ReadBinaryWatch(num int) []string {
	var res []string = make([]string,0)
	if num <= 0{
		res = append(res, "0:00")
		return res
	}

	var times []int
	combine(times,&res,0,num,0,10)
	return res
}

func countbit(n int)int{
	var cnt int = 0
	for n != 0{
		n = n & (n - 1)
		cnt++
	}
	return cnt
}
func ReadBinaryWatch2(num int) []string{
	var res []string
	for i := 0;i < 12;i++{
		for j := 0;j < 60;j++{
			if countbit(i) + countbit(j) == num{
				m := fmt.Sprintf("%.2d",j)
				res = append(res,strconv.Itoa(i) + ":" + m)
			}
		}
	}
	return res
}