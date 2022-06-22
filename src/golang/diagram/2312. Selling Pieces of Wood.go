package diagram

func dp_sellingWood(rows int,columns int,prices [][]int,memo [][]int64)int64{
	if rows <= 0 || columns <= 0{
		return 0
	}
	if memo[rows][columns] != -1{
		return int64(memo[rows][columns])
	}
	var res int64 = int64(prices[rows][columns])
	for i := 1;i < 1 + rows/2;i++{
		cur := dp_sellingWood(rows - i,columns,prices,memo) + dp_sellingWood(i,columns,prices,memo)
		res = max_int64(res,cur)
	}
	for j := 1;j < 1 + columns/2;j++{
		cur := dp_sellingWood(rows,columns - j,prices,memo) + dp_sellingWood(rows,j,prices,memo)
		res = max_int64(res,cur)
	}
	//for _,p := range prices{
	//	h := p[0]
	//	w := p[1]
	//	if h <= rows && w <= columns{
	//		cur1 := int64(p[2]) + dp_sellingWood(rows - h,columns,prices,memo) + dp_sellingWood(h,columns - int(p[1]),prices,memo)
	//		res = max_int64(res,cur1)
	//		cur2 := int64(p[2]) + dp_sellingWood(rows,columns - w,prices,memo) + dp_sellingWood(rows - int(p[0]),w,prices,memo)
	//		res = max_int64(res,cur2)
	//	}
	//}
	memo[rows][columns] = res
	return res
}

func SellingWood(m int, n int, prices [][]int) int64 {
	var p [][]int = make([][]int,m + 1)
	for i := 0;i <= m;i++{
		p[i] = make([]int,n + 1)
	}
	for _,price := range prices{
		p[price[0]][price[1]] = int(price[2])
	}
	var memo [][]int64 = make([][]int64,m + 1)
	for i := 0;i <= m;i++{
		memo[i] = make([]int64,n + 1)
		for j := 0;j <= n;j++{
			memo[i][j] = -1
		}
	}
	return dp_sellingWood(m,n,p,memo)
}