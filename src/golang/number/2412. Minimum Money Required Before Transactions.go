package number

import "sort"

func minimumMoney(transactions [][]int) int64{
	var max_need int64 = 0
	var spend int64 = 0
	for _,trans := range transactions{
		if trans[0] > trans[1]{
			spend += int64(trans[0] - trans[1])
			max_need = max_int64(max_need,int64(trans[1]))
		}else{
			max_need = max_int64(max_need,int64(trans[0]))
		}
	}
	return spend + max_need
}

func MinimumMoney(transactions [][]int) int64 {
	var l int = len(transactions)
	var earn [][]int
	var neg [][]int
	for i := 0;i < l;i++{
		if transactions[i][1] >= transactions[i][0]{
			earn = append(earn,[]int{transactions[i][0],transactions[i][1]})
		}else{
			neg = append(neg,[]int{transactions[i][0],transactions[i][1]})
		}
	}
	sort.Slice(earn, func(i, j int) bool {
		return earn[i][0] > earn[j][0]
	})
	sort.Slice(neg, func(i, j int) bool {
		if neg[i][1] == neg[j][1]{
			return neg[i][0] < neg[j][0]
		}
		return neg[i][1] < neg[j][1]
	})
	var cur int64 = 0
	var res int64 = 0
	for i := 0;i < len(neg);i++{
		if cur < int64(neg[i][0]){
			res += int64(neg[i][0]) - cur
			cur = int64(neg[i][1])
		}else{
			cur += int64(neg[i][1] - neg[i][0])
		}
	}
	for i := 0;i < len(earn);i++{
		if cur < int64(earn[i][0]){
			res += int64(earn[i][0]) - cur
			cur = int64(earn[i][1])
		}else{
			cur += int64(earn[i][1] - earn[i][0])
		}
	}
	return res
}