package number

//需要将 所有商品 分配到零售商店
//每个店最多分配mid个商品，
//分配后，每间商店都会被分配一定数目的商品（可能为 0 件）。用 x 表示所有商店中分配商品数目的最大值，
//你希望 x 越小越好。也就是说，你想 最小化 分配给任意商店商品数目的 最大值

//假设最大值为mid，能否将所有商品分配完？ 能分配完返回true，不然返回true
func check_minimizedMaximum(n int,quantities []int,mid int)bool{
	var l int = len(quantities)
	for i := 0;i < l;i++{
		n -= quantities[i] / mid
		if quantities[i] % mid != 0{
			n--
		}
	}
	if n >= 0{
		return true
	}
	return false
}

func MinimizedMaximum(n int, quantities []int) int {
	var low int = 1
	var high int = 0
	var l int = len(quantities)
	for i := 0;i < l;i++{
		if quantities[i] > high{
			high = quantities[i]
		}
	}
	var res int = low
	for low < high{
		mid := (low + high)/2
		ret := check_minimizedMaximum(n,quantities,mid)
		if ret{
			res = mid
			high = mid
		}else{
			low = mid + 1
			res = mid + 1
		}
	}
	return res
}