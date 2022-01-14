package string_issue

func WordCount(startWords []string, targetWords []string) int {
	var record map[int]bool = make(map[int]bool)
	for _, word := range startWords {
		var mask int = 0
		for _, c := range word {
			mask |= 1 << (c - 'a')
		}
		record[mask] = true
	}
	var res int = 0
	for _, word := range targetWords {
		var cur_len int = len(word)
		var mask int = 0
		for i := 0; i < cur_len; i++ {
			mask |= 1 << (word[i] - 'a')
		}
		for i := 0; i < cur_len; i++ {
			cur := mask ^ (1 << (word[i] - 'a'))
			if _, ok := record[cur]; ok {
				res++
				break
			}
		}
	}
	return res
}

//func WordCount(startWords []string, targetWords []string) int {
//	var record map[string]int = make(map[string]int)
//	for _, word := range targetWords {
//		var data []uint8
//		for i := 0; i < len(word); i++ {
//			data = append(data, word[i])
//		}
//		sort.Slice(data, func(i, j int) bool {
//			return data[i] < data[j]
//		})
//		var sb strings.Builder
//		for _, d := range data {
//			sb.WriteByte(d)
//		}
//		record[sb.String()]++
//	}
//	var res int = 0
//	for _, word := range startWords {
//		var data []uint8
//		for i := 0; i < len(word); i++ {
//			data = append(data, word[i])
//		}
//		sort.Slice(data, func(i, j int) bool {
//			return data[i] < data[j]
//		})
//		var sb strings.Builder
//		for _, d := range data {
//			sb.WriteByte(d)
//		}
//		var l1 int = len(data)
//		for k, v := range record {
//			var l2 int = len(k)
//			if l1+1 != l2 {
//				continue
//			}
//			var idx1 int = 0
//			var idx2 int = 0
//			for idx1 < l1 && idx2 < l2 {
//				if data[idx1] == k[idx2] {
//					idx1++
//					idx2++
//				} else {
//					idx2++
//				}
//			}
//			if idx1 == l1 && (idx2 == l2 || idx2 == l2-1) {
//				res += v
//				delete(record, k)
//			}
//		}
//	}
//	return res
//}

//type Tie struct {
//	next    [26]*Tie
//	is_word bool
//}
//
//func WordCount(startWords []string, targetWords []string) int {
//	var root Tie
//	for _, word := range startWords {
//		var data []uint8
//		for i := 0; i < len(word); i++ {
//			data = append(data, word[i]-'a')
//		}
//		sort.Slice(data, func(i, j int) bool {
//			return word[i] < word[j]
//		})
//		visit := &root
//		for _, d := range data {
//			if visit.next[d] == nil {
//				visit.next[d] = new(Tie)
//			}
//			visit = visit.next[d]
//		}
//		visit.is_word = true
//	}
//	var res int = 0
//	for _, word := range targetWords {
//		var data []uint8
//		for i := 0; i < len(word); i++ {
//			data = append(data, word[i]-'a')
//		}
//		sort.Slice(data, func(i, j int) bool {
//			return data[i] < data[j]
//		})
//		visit := &root
//		var wrong int = 0
//		for _, d := range data {
//			if visit.next[d] == nil {
//				wrong++
//				if wrong > 1 {
//					break
//				}
//			} else {
//				visit = visit.next[d]
//			}
//		}
//		if wrong == 1 && visit.is_word {
//			res++
//		}
//	}
//	return res
//}
