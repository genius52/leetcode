package string_issue

//S1 = "0"
//S2 = "011"
//S3 = "0111001"
//S4 = "011100110110001"

//输入：n = 3, k = 1
//输出："0"
//解释：S3 为 "0111001"，其第 1 位为 "0" 。
func recursive_findKthBit(n int,k int)byte{
	cur_len := (1 << n) - 1
	mid := cur_len/2 + 1
	if n == 1{
		return '0'
	}
	if k == mid{
		return '1'
	}else if k < mid{
		return recursive_findKthBit(n - 1,k)
	}else{
		ret := recursive_findKthBit(n - 1,cur_len - k + 1)
		if ret == '0'{
			return '1'
		}else{
			return '0'
		}
	}
}

func FindKthBit(n int, k int) byte{
	return recursive_findKthBit(n,k)
}

//func reverse_invert(record []uint8) []uint8{
//	var reverse []uint8
//	var l int = len(record)
//	for i := l - 1;i >= 0;i--{
//		if record[i] == '0'{
//			reverse = append(reverse,'1')
//		}else{
//			reverse = append(reverse,'0')
//		}
//	}
//	record = append(record,'1')
//	record = append(record,reverse...)
//	return record
//}
//
//func FindKthBit(n int, k int) byte {
//	var s_len int = 1
//	var record []uint8 = []uint8{'0'}
//	for i := 1;i < n;i++{
//		record = reverse_invert(record)
//		s_len = s_len * 2 + 1
//		if(k <= s_len){
//			return record[k - 1]
//		}
//	}
//	return '0'
//}