package number

func ProductQueries(n int, queries [][]int) []int{
	var l int = len(queries)
	var res []int = make([]int,l)
	var data []int
	var MOD int64 = 1e9 + 7
	for offset := 0;offset < 32;offset++{
		if n | (1 << offset) == n{
			data = append(data,offset)
		}
	}
	for i := 0;i < l;i++{
		left := queries[i][0]
		right := queries[i][1]
		res[i] = 1
		for j := left;j <= right;j++{
			res[i] = int(int64(res[i]) * int64(1 << data[j]) % MOD)
		}
	}
	return res
}

//func ProductQueries(n int, queries [][]int) []int {
//	var l int = len(queries)
//	var res []int = make([]int,l)
//	var data []int
//	var MOD int = 1e9 + 7
//	for n > 0{
//		p := int(math.Log2(float64(n)))
//		n2 := int(math.Pow(2,float64(p)))
//		data = append(data,n2)
//		n -= n2
//	}
//	sort.Ints(data)
//
//	var l2 int = len(data)
//	var prefix []int = make([]int,l2 + 1)
//	prefix[0] = 1
//	for i := 0;i < l2;i++{
//		prefix[i + 1] = prefix[i] * data[i]
//		prefix[i + 1] %= MOD
//	}
//	for i := 0;i < l;i++{
//		res[i] = prefix[queries[i][1] + 1] / prefix[queries[i][0]] % MOD
//	}
//	return res
//}