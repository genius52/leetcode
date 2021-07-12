package number

//Alice starting first.
//For Bob to win, the sum of the digits in the first half of num must be equal to the sum of the digits in the second half.
//For Alice to win, the sums must not be equal.
//Assuming Alice and Bob play optimally, return true if Alice will win and false if Bob will win.
func SumGame(num string) bool {
	var tag1 int = 0
	var tag2 int = 0
	var sum1 int = 0
	var sum2 int = 0
	var l int = len(num)
	for i := 0;i < l;i++{
		if num[i] == '?'{
			if i < l/2{
				tag1++
			}else{
				tag2++
			}
		}else{
			if i < l/2{
				sum1 += int(num[i] - '0')
			}else{
				sum2 += int(num[i] - '0')
			}
		}
	}
	if tag1 == 0 && tag2 == 0{
		return sum1 != sum2
	}
	if (tag1 + tag2) % 2 == 1{
		return true
	}

	if sum1 > sum2{
		if tag1 >= tag2{
			return true
		}
	}else if sum1 < sum2{
		if tag1 <= tag2{
			return true
		}
	}else{
		return false
	}
	return false
}