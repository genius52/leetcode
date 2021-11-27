package array

//Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
//Output: 0
//students = [1,1,1,0,0,1], sandwiches = [1,0,0,0,1,1]
func CountStudents2(students []int, sandwiches []int) int {
	var zero_students int = 0
	var one_students int = 0
	for _,n := range students{
		if n == 0{
			zero_students++
		}else{
			one_students++
		}
	}
	var l int = len(sandwiches)
	var pos int = 0
	for pos < l{
		if sandwiches[pos] == 0{
			if zero_students == 0{
				break
			}
			zero_students--
		}else{
			if one_students == 0{
				break
			}
			one_students--
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