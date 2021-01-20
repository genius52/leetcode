package array

//Input: n = 5
//Output: [5,3,1,4,3,5,2,4,2]
func dfs_constructDistancedSequence(max_num int,data []int,index int,l int,visited []bool)bool{
	if index >= l{
		for i := 1;i < max_num;i++{
			if !visited[i]{
				return false
			}
		}
		return true
	}
	if data[index] != 0{
		return dfs_constructDistancedSequence(max_num,data,index + 1,l,visited)
	}
	for i := max_num;i >= 1;i--{
		if visited[i]{
			continue
		}
		if i != 1 && (index + i >= l || data[index + i] != 0){
			continue
		}
		data[index] = i
		if i != 1{
			data[index + i] = i
		}
		visited[i] = true
		ret := dfs_constructDistancedSequence(max_num,data,index + 1,l,visited)
		if !ret{
			data[index] = 0
			if i != 1{
				data[index + i] = 0
			}
			visited[i] = false
		}else{
			return true
		}
	}
	return false
}

func ConstructDistancedSequence(n int) []int {
	var l int = (n - 1) * 2 + 1
	var data []int = make([]int,l)
	var visited []bool = make([]bool,n + 1)
	dfs_constructDistancedSequence(n,data,0,l,visited)
	return data
}