package number

import "sort"

// restaurants[i] = [idi, ratingi, veganFriendlyi, pricei, distancei]
//Input: restaurants = [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]],
//veganFriendly = 1, maxPrice = 50, maxDistance = 10
//Output: [3,1,5]
//按照 rating 从高到低排序。如果 rating 相同，那么按 id 从高到低排序。
func FilterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var l int = len(restaurants)
	var idx []int
	for i := 0;i < l;i++{
		if (veganFriendly == 1 && restaurants[i][2] == 0) || restaurants[i][3] > maxPrice || restaurants[i][4] > maxDistance{
			continue
		}
		idx = append(idx,i)
	}
	sort.Slice(idx, func(i, j int) bool {
		if restaurants[idx[i]][1] == restaurants[idx[j]][1]{
			return restaurants[idx[i]][0] > restaurants[idx[j]][0]
		}
		r1 := restaurants[idx[i]][1]
		r2 := restaurants[idx[j]][1]
		return  r1 > r2
	})
	var idx_len int = len(idx)
	var res []int = make([]int,idx_len)
	for i := 0;i < idx_len;i++{
		r := restaurants[idx[i]]
		res[i] = r[0]
	}
	return res
}