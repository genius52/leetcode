package tree

//Input: nums = [3,4,5,1,2]
//Output: 5
//Explanation: The following 5 arrays will yield the same BST:
//[3,1,2,4,5]
//[3,1,4,2,5]
//[3,1,4,5,2]
//[3,4,1,2,5]
//[3,4,1,5,2]
func combine_numOfWays(n int,m int)int{
	return 0
}

func dfs_NumOfWays(nums []int)int{
	var l int = len(nums)
	if l <= 2{
		return 1
	}
	var less []int
	var greater []int
	for i := 1;i < l;i++{
		if nums[i] < nums[0]{
			less = append(less,nums[i])
		}
		if nums[i] > nums[0]{
			greater = append(greater,nums[i])
		}
	}
	if len(less) == 0 || len(greater) == 0{
		return 1
	}
	//组合数 * 小于nums[i]的全排列 * 大于nums[i]的全排列
	//l - 1个坑，选出其中的len(less)个坑放小于nums[0]的数
	return combine_numOfWays(l - 1, len(less)) * dfs_NumOfWays(less) * dfs_NumOfWays(greater)
}

func NumOfWays(nums []int) int{
	return dfs_NumOfWays(nums)
}