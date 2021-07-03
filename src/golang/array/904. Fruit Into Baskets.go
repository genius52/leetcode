package array

//Input: fruits = [3,3,3,1,2,1,1,2,3,3,4]
//Output: 5
//Explanation: We can pick from trees [1,2,1,1,2].
func TotalFruit(tree []int) int {
	var l int = len(tree)
	var left int = 0
	var right int = 0
	var res int = 0
	var cnt map[int]int = make(map[int]int)
	for left < l{
		for right < l{
			if _,ok := cnt[tree[right]];ok{
				cnt[tree[right]]++
			}else{
				cnt[tree[right]] = 1
			}
			if len(cnt) > 2{
				break
			}
			cur := right - left + 1
			if cur > res{
				res = cur
			}
			right++
		}
		for left < right{
			cnt[tree[left]]--
			if cnt[tree[left]] == 0{
				delete(cnt,tree[left])
			}
			left++
			if len(cnt) <= 2{
				break
			}
		}
		right++
	}
	return res
}