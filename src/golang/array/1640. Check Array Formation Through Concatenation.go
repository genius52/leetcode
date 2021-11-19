package array

func dfs_canFormArray(arr []int,l1 int,pos1 int,pieces [][]int,l2 int)bool{
	if pos1 >= l1{
		return true
	}
	for i := 0;i < l2;i++{
		var match bool = true
		for j := 0;j < len(pieces[i]);j++{
			if arr[pos1 + j] != pieces[i][j]{
				match = false
				break
			}
		}
		if match{
			ret := dfs_canFormArray(arr,l1,pos1 + len(pieces[i]),pieces,l2)
			if ret{
				return true
			}
		}
	}
	return false
}

func canFormArray(arr []int, pieces [][]int) bool{
	return dfs_canFormArray(arr,len(arr),0,pieces,len(pieces))
}

func CanFormArray(arr []int, pieces [][]int) bool {
	var l1 int = len(arr)
	var l2 int = len(pieces)
	var index int = 0
	for index < l1{
		var find bool = false
		for j := 0;j < l2;j++{
			var k int = 0
			for k < len(pieces[j]){
				if pieces[j][k] == 0{
					k++
					continue
				}
				if pieces[j][k] == arr[index]{
					find = true
					pieces[j][k] = 0
				}
				break
			}
			if find{
				break
			}
		}
		if !find{
			return false
		}
		index++
	}
	return true
}