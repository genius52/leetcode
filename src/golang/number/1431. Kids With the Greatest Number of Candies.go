package number

//Input: candies = [2,3,5,1,3], extraCandies = 3
//Output: [true,true,true,false,true]
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int = 0
	for _,v := range candies{
		if v > max{
			max = v
		}
	}
	var res []bool = make([]bool,len(candies))
	for i,v := range candies{
		res[i] = (v + extraCandies) >= max
	}
	return res
}
