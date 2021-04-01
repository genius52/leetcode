package number

//Input:
//n: 13   k: 2
//Output:
//10
//Explanation:
//The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.

func dfs_findKthNumber(n int,k *int,cur int)int{
	if cur > n{
		return -1
	}
	if *k == 0{
		return cur
	}
	for i := 0;i <= 9;i++{
		if cur == 0 && i == 0{
			continue
		}
		next := cur * 10 + i
		if next > n{
			continue
		}

		*k--
		res := dfs_findKthNumber(n, k,next)
		if res != -1{
			return res
		}
	}
	return -1
}

func FindKthNumber(n int, k int) int {
	return dfs_findKthNumber(n,&k,0)
}