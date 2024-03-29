package array

import "sort"

//Input: A = [1,1,2,2,3,3,4,4,5,5], target = 8
//Output: 20
//Explanation:
//Enumerating by the values (A[i], A[j], A[k]):
//(1, 2, 5) occurs 8 times;
//(1, 3, 4) occurs 8 times;
//(2, 2, 4) occurs 2 times;
//(2, 3, 3) occurs 2 times.
func ThreeSumMulti(A []int, target int) int{
	var l int = len(A)
	var sum_cnt map[int]int = make(map[int]int)
	var num_cnt map[int]int = make(map[int]int)
	var res int = 0
	for i := 0;i < l;i++{
		cur_target := target - A[i]
		if cnt,ok := sum_cnt[cur_target];ok{
			res += cnt
			res %= 1000000007
		}
		for n,cnt := range num_cnt{
			if _,ok := sum_cnt[n + A[i]];ok{
				sum_cnt[n + A[i]] += cnt
			}else{
				sum_cnt[n + A[i]] = cnt
			}
		}
		num_cnt[A[i]]++
	}
	return res % 1000000007
}

func ThreeSumMulti2(arr []int, target int) int{
	sort.Ints(arr)
	var l int = len(arr)
	var res int = 0
	var MOD int = 1e9 + 7
	for left := 0;left < l - 2;left++{
		var mid int = left + 1
		var right int = l - 1
		for mid < right{
			cur_sum := arr[left] + arr[mid] + arr[right]
			if cur_sum == target{
				if arr[mid] != arr[right]{
					cnt1 := 1
					for (mid + 1) < right && (arr[mid] == arr[mid + 1]){
						mid++
						cnt1++
					}
					cnt2 := 1
					for mid < (right - 1) && (arr[right - 1] == arr[right]){
						right--
						cnt2++
					}
					res += cnt1 * cnt2
					res %= MOD
					mid++
					right--
				}else{
					res += ((right - mid + 1) * (right - mid)) /2
					res %= MOD
					break
				}
			}else if cur_sum < target{
				mid++
			}else if cur_sum > target{
				right--
			}
		}
	}
	return res
}

//func ThreeSumMulti(A []int, target int) int{
//	var record map[int]int = make(map[int]int)
//	var res int = 0
//	var l int = len(A)
//	for i := 0;i < l;i++{
//		cur_tar := target - A[i]
//		if cnt,ok := record[cur_tar];ok{
//			res += cnt
//			res = res % 1000000007
//		}
//		for j := 0;j < i;j++{
//			if i == j{
//				continue
//			}
//			cur_sum := A[i] + A[j]
//			if _,ok := record[cur_sum];ok{
//				record[cur_sum]++
//			}else{
//				record[cur_sum] = 1
//			}
//		}
//	}
//	return res
//}

//func ThreeSumMulti(A []int, target int) int {
//	var record map[int]int = make(map[int]int)
//	for _,a := range A{
//		record[a]++
//	}
//	var nums []int
//	for n,_ := range record{
//		nums = append(nums,n)
//	}
//	sort.Ints(nums)
//	var res int = 0
//	var l int = len(nums)
//	for i := 0;i < l;i++{
//		if record[nums[i]] >= 3{
//			if nums[i] * 3 == target{
//				res = record[nums[i]] * (record[nums[i]] - 1) * (record[nums[i]] - 2)/6
//				res = res % 1000000007
//			}
//		} else if record[nums[i]] == 2{
//			var target1 int = target - nums[i] * 2
//			if target1 != nums[i]{
//				if cnt,ok := record[target1];ok{
//					res += cnt
//					res = res % 1000000007
//				}
//			}
//		}
//		for j := i + 1;j < l;j++{
//			var target2 int = target - nums[i] - nums[j]
//			if target2 <= nums[j] {
//				break
//			}
//			if cnt,ok := record[target2];ok{
//				res += record[nums[i]] * record[nums[j]] * cnt
//				res = res % 1000000007
//			}
//		}
//	}
//	return res
//}