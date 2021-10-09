package number

//Input: s = "1101"
//Output: 6
//Explanation: "1101" corressponds to number 13 in their decimal representation.
//Step 1) 13 is odd, add 1 and obtain 14.
//Step 2) 14 is even, divide by 2 and obtain 7.
//Step 3) 7 is odd, add 1 and obtain 8.
//Step 4) 8 is even, divide by 2 and obtain 4.
//Step 5) 4 is even, divide by 2 and obtain 2.
//Step 6) 2 is even, divide by 2 and obtain 1.
func NumSteps(s string) int {
	var l int = len(s)
	var steps int = 0
	var upgrade bool = false
	for i := l - 1;i >= 1;i--{
		steps++
		if s[i] == '1'{
			if !upgrade{
				steps++
			}
			upgrade = true
		}else{
			if upgrade{
				steps++
			}
		}
	}
	if upgrade{
		steps++
	}
	return steps
}