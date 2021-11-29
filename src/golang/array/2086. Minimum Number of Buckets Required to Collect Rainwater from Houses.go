package array

import "strings"

func minimumBuckets(street string) int {
	var l int = len(street)
	if l == 1{
		if street[0] == 'H'{
			return -1
		}else{
			return 0
		}
	}
	if strings.Contains(street,"HHH") || strings.HasPrefix(street,"HH") || strings.HasSuffix(street,"HH"){
		return -1
	}
	var res int = 0
	var data []byte = []byte(street)
	for i := 1;i < l - 1;i++{
		if data[i] == '.' && data[i - 1] == 'H' && data[i + 1] == 'H'{
			data[i - 1] = '.'
			data[i + 1] = '.'
			res++
		}
	}
	for i := 0;i < l;i++{
		if data[i] == 'H'{
			res++
		}
	}
	return res
}