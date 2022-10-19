package diagram

func bfs_componentValue(nums []int,leaves []int,graph [][]int,degrees []int,target int)bool{
	for len(leaves) > 0{
		var leaves2 []int
		var cur_len int = len(leaves)
		for i := 0;i < cur_len;i++{
			if nums[leaves[i]] > target{
				return false
			}
			for _,next := range graph[leaves[i]]{
				if nums[leaves[i]] < target{
					nums[next] += nums[leaves[i]]
				}
				degrees[next]--
				if degrees[next] == 1{
					leaves2 = append(leaves2,next)
				}
			}
		}
		leaves = leaves2
	}
	return true
}

func ComponentValue(nums []int, edges [][]int) int {
	var sum int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	var graph [][]int = make([][]int,l)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var leaves []int
	var degrees []int = make([]int,l)
	for i := 0;i < l;i++{
		degrees[i] = len(graph[i])
		if len(graph[i]) == 1{
			leaves = append(leaves,i)
			//q.PushBack(i)
		}
	}

	for i := l;i >= 1;i--{
		if sum % i == 0{
			target := sum / i
			var nums2 []int = make([]int,l)
			copy(nums2,nums)
			var leaves2 []int = make([]int,len(leaves))
			copy(leaves2,leaves)
			var degrees2 []int = make([]int,l)
			copy(degrees2,degrees)
			if bfs_componentValue(nums2,leaves2,graph,degrees2,target){
				return i - 1
			}
		}
	}
	return 0
}