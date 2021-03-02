package string_issue

import "strconv"

func IsIsomorphic(s string, t string) bool{
	var len1 int = len(s)
	var len2 int = len(t)
	if len1 != len2{
		return false
	}
	var count1 [256]int
	var count2 [256]int
	for i := 0;i < len1;i++{
		if count1[s[i]] != count2[t[i]]{ //check the last occur position
			return false
		}
		count1[s[i]] = i + 1
		count2[t[i]] = i + 1
		if count1[s[i]] != count2[t[i]]{
			return false
		}
	}
	return true
}

func IsIsomorphic2(s string, t string) bool {
	var len1 int = len(s)
	var len2 int = len(t)
	if len1 != len2{
		return false
	}
	var pos1 map[uint8][]int = make(map[uint8][]int)
	var pos2 map[uint8][]int = make(map[uint8][]int)
	for i := 0;i < len1;i++{
		pos1[s[i]] = append(pos1[s[i]],i)
		pos2[t[i]] = append(pos2[t[i]],i)
	}
	var record1 map[string]bool = make(map[string]bool)
	var record2 map[string]bool = make(map[string]bool)
	for _,val := range pos1{
		var k string
		for _,p := range val{
			if len(k) != 0{
				k += ","
			}
			k += strconv.Itoa(p)
		}
		record1[k] = true
	}
	for _,val := range pos2{
		var k string
		for _,p := range val{
			if len(k) != 0{
				k += ","
			}
			k += strconv.Itoa(p)
		}
		record2[k] = true
	}
	for k,_ := range record1{
		if _,ok := record2[k];!ok{
			return false
		}
	}
	return true
}