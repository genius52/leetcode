package string_issue

import "strconv"

func OriginalDigits(s string) string {
	var record []int = make([]int,10)
	for _,c := range s{
		if(c == 'z'){
			record[0]++
		}else if(c == 'w'){
			record[2]++
		}else if(c == 'u'){
			record[4]++
		}else if(c == 'x'){
			record[6]++
		}else if(c == 'g'){
			record[8]++
		}else if(c == 'o'){
			record[1]++ //one = record[1] - record[2] - record[4] - record[0]
		}else if(c == 't'){
			record[3]++//three = record[3] - record[2] - record[8]
		}else if(c == 'f'){
			record[5]++//five = record[5] - record[4]
		}else if(c == 's'){
			record[7]++//sever = record[7] - record[6]
		}else if(c == 'i'){
			record[9]++//nine = record[9] - record[8] - record[5] - record[6]
		}
	}
	var res string
	record[1] = record[1] - record[2] - record[4] - record[0]
	record[3] = record[3] - record[2] - record[8]
	record[5] = record[5] - record[4]
	record[7] = record[7] - record[6]
	record[9] = record[9] - record[8] - record[5] - record[6]
	for i := 0;i <= 9;i++{
		for record[i] > 0{
			res += strconv.Itoa(i)
			record[i]--
		}
	}
	return res
}
