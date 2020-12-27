package array

//Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
//Output: 0
func CountStudents2(students []int, sandwiches []int) int {
	var zero_cnt int = 0
	var one_cnt int = 0
	for _,n := range students{
		if n == 0{
			zero_cnt++
		}else{
			one_cnt++
		}
	}
	var l int = len(sandwiches)
	var pos int = 0
	for pos < l{
		if sandwiches[pos] == 0{
			if zero_cnt == 0{
				break
			}
			zero_cnt--
		}else{
			if one_cnt == 0{
				break
			}
			one_cnt--
		}
		pos++
	}
	return l - pos
}

func CountStudents(students []int, sandwiches []int) int {
	for len(sandwiches) > 0 && len(students) > 0{
		sand := sandwiches[0]
		sandwiches = sandwiches[1:]
		var stu_len int = len(students)
		pos := 0
		for pos < stu_len{
			if students[pos] == sand{
				break
			}
			pos++
		}
		if pos < stu_len{
			students = append(students[pos + 1:],students[:pos]...)
		}else{
			return stu_len
		}
	}
	return 0
}