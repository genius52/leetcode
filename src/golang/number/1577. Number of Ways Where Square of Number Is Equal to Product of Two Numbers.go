package number

import "strconv"

//Input: nums1 = [7,7,8,3], nums2 = [1,2,9,7]
//Output: 2
//Explanation: There are 2 valid triplets.
//Type 1: (3,0,2).  nums1[3]^2 = nums2[0] * nums2[2].
//Type 2: (3,0,1).  nums2[3]^2 = nums1[0] * nums1[1].
func NumTriplets(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var product_record1 map[int]int = make(map[int]int)
	var product_record2 map[int]int = make(map[int]int)
	var num_record1 map[int]int = make(map[int]int)
	var num_record2 map[int]int = make(map[int]int)
	for i := 0;i < l1;i++{
		val := nums1[i] * nums1[i]
		if _,ok := product_record1[val];ok{
			product_record1[val]++
		}else{
			product_record1[val] = 1
		}
		if _,ok := num_record1[nums1[i]];ok{
			num_record1[nums1[i]]++
		}else{
			num_record1[nums1[i]] = 1
		}
	}
	for i := 0;i < l2;i++{
		val := nums2[i] * nums2[i]
		if _,ok := product_record2[val];ok{
			product_record2[val]++
		}else{
			product_record2[val] = 1
		}
		if _,ok := num_record2[nums2[i]];ok{
			num_record2[nums2[i]]++
		}else{
			num_record2[nums2[i]] = 1
		}
	}
	var res int = 0
	for product,product_cnt := range product_record1{
		var visited map[string]bool = make(map[string]bool)
		for num,num_cnt := range num_record2{
			if product % num != 0{
				continue
			}
			need_num := product / num
			if need_num == num{
				res += product_cnt * num_cnt * (num_cnt - 1)/2
			}else{
				if need_num_cnt,ok := num_record2[need_num];!ok{
					continue
				}else{
					k := strconv.Itoa(num) + "," + strconv.Itoa(need_num)
					if _,ok := visited[k];ok{
						continue
					}
					res += product_cnt * num_cnt * need_num_cnt
					visited[k] = true
					visited[strconv.Itoa(need_num) + "," + strconv.Itoa(num) ] = true
				}
			}
		}
	}

	for product,product_cnt := range product_record2{
		var visited map[string]bool = make(map[string]bool)
		for num,num_cnt := range num_record1{
			if product % num != 0{
				continue
			}
			need_num := product / num
			if need_num == num{
				res += product_cnt * num_cnt * (num_cnt - 1)/2
			}else{
				if need_num_cnt,ok := num_record1[need_num];!ok{
					continue
				}else{
					k := strconv.Itoa(num) + "," + strconv.Itoa(need_num)
					if _,ok := visited[k];ok{
						continue
					}
					res += product_cnt * num_cnt * need_num_cnt
					visited[k] = true
					visited[strconv.Itoa(need_num) + "," + strconv.Itoa(num) ] = true
				}
			}
		}
	}
	return res
}