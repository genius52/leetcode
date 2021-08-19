package number

//分别返回样本的最小值、最大值、平均值、中位数和众数。其中，众数是保证唯一的
func SampleStats(count []int) []float64 {
	var res []float64 = make([]float64,5)
	res[0] = 255
	res[1] = 0
	var frequency int = 0
	var sum float64 = 0
	var cnt int = 0
	for i := 0;i <= 255;i++{
		if count[i] == 0{
			continue
		}
		if count[i] > frequency{
			frequency = count[i]
			res[4] = float64(i)
		}
		if float64(i) < res[0]{
			res[0] = float64(i)
		}
		if float64(i) > res[1]{
			res[1] = float64(i)
		}
		sum += float64(i) * float64(count[i])
		cnt += count[i]
	}
	res[2] = sum / float64(cnt)
	var first int = 0
	var second int = 0
	var num1 int = -1
	var num2 int = -1
	if (cnt | 1) == cnt{
		first = cnt/2 + 1
		second = cnt/2 + 1
	}else{
		first = cnt/2
		second = (cnt + 1)/2 + 1
	}
	for i := 0;i <= 255;i++{
		if count[i] == 0{
			continue
		}
		if first > 0{
			first -= count[i]
		}
		if second > 0{
			second -= count[i]
		}

		if num1 == -1 && first <= 0{
			num1 = i
		}
		if num2 == -1 && second <= 0{
			num2 = i
			break
		}
	}
	res[3] = float64(num1 + num2)/2
	return res
}