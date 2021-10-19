package array

func dp_maxDotProduct(nums1 []int, nums2 []int,l1 int,l2 int,pos1 int,pos2 int,choose bool,memo *[][][2]int)int{
	if pos1 == l1 || pos2 == l2{
		if choose{
			return 0
		}
		return -500000001
	}
	var tag int = 0
	if choose{
		tag = 1
	}
	if (*memo)[pos1][pos2][tag] != -2147483648{
		return (*memo)[pos1][pos2][tag]
	}
	//skip current or choose pos1 && pos2
	ret1 := nums1[pos1] * nums2[pos2] + dp_maxDotProduct(nums1,nums2,l1,l2,pos1 + 1,pos2 + 1,true,memo)
	//choose pos1,skip pos2
	ret2 := dp_maxDotProduct(nums1,nums2,l1,l2,pos1,pos2 + 1,choose,memo)
	//skip pos1,choose pos2
	ret3 := dp_maxDotProduct(nums1,nums2,l1,l2,pos1 + 1,pos2,choose,memo)
	(*memo)[pos1][pos2][tag] = max_int_number(ret1,ret2,ret3)
	return (*memo)[pos1][pos2][tag]
}

func MaxDotProduct(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var memo [][][2]int = make([][][2]int,l1)
	for i := 0;i < l1;i++{
		memo[i] = make([][2]int,l2)
		for j := 0;j < l2;j++{
			memo[i][j][0] = -2147483648
			memo[i][j][1] = -2147483648
		}
	}
	return dp_maxDotProduct(nums1,nums2,l1,l2,0,0,false,&memo)
}

func MaxDotProduct2(nums1 []int, nums2 []int) int{
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var dp [][]int = make([][]int,l1 + 1)
	for i := 0;i <= l1;i++{
		dp[i] = make([]int,l2 + 1)
		for j := 0;j <= l2;j++{
			dp[i][j] = -2147483648
		}
	}
	for i := 1;i <= l1;i++{
		for j := 1;j <= l2;j++{
			dp[i][j] = max_int_number(nums1[i - 1] * nums2[j - 1] + dp[i - 1][j - 1],
				dp[i - 1][j],
				dp[i][j - 1],
				nums1[i - 1] * nums2[j - 1])
		}
	}
	return dp[l1][l2]
}