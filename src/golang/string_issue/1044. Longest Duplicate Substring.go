package string_issue

//robin karp
func check_longestDupSubstring(s string,sub_len int,power *[]int64)string{
	var record map[int64][]string = make(map[int64][]string)
	var l int = len(s)
	var hash_val int64 = 0
	for i := 0;i < sub_len;i++{
		hash_val = hash_val * 26 % 19260817 + int64(s[i] - 'a')
		hash_val = hash_val % 19260817
	}
	record[hash_val] = append(record[hash_val],s[:sub_len])
	for i := 0;i + sub_len < l;i++{
		//减去第一个，加上新的一个
		hash_val = ((hash_val - (*power)[sub_len - 1] * int64(s[i] - 'a')) % 19260817 + 19260817) % 19260817
		hash_val = (hash_val * 26 + int64(s[i + sub_len] - 'a')) % 19260817
		if pre,ok := record[hash_val];ok{
			for _,p := range pre{
				if p == s[i + 1:i + 1 + sub_len]{
					return p
				}
			}
		}
		record[hash_val] = append(record[hash_val],s[i + 1:i + 1 + sub_len])
	}
	return ""
}

func LongestDupSubstring(s string) string {
	var l int = len(s)
	var low int = 1
	var high int = l - 1
	var power []int64 = make([]int64,l)
	power[0] = 1
	for i := 1;i < l;i++{
		power[i] = (power[i - 1] * 26) % 19260817
	}
	var res string
	for low <= high{
		mid := low + (high - low)/2
		cur := check_longestDupSubstring(s,mid,&power)
		if len(cur) > 0{
			res = cur
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return res
}