package number

func distributeCandies(candies int, num_people int) []int {
	var res []int = make([]int,num_people,num_people)
	var cnt int = 1
	for candies > cnt{
		res[(cnt - 1) % num_people] += cnt
		candies -= cnt
		cnt++
	}
	res[(cnt - 1) % num_people] += candies
	return res
}

func distributeCandies2(candies int, num_people int) []int {
	var res []int = make([]int,num_people)
	var candy_num int = 1
	var idx int = 0
	for candies > 0{
		if candies >= candy_num{
			res[idx] += candy_num
		}else{
			res[idx] += candies
		}
		candies -= candy_num
		candy_num++
		idx = (idx + 1) % num_people
	}
	return res
}