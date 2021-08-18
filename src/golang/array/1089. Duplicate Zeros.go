package array

func duplicateZeros(arr []int)  {
	var l int = len(arr)
	var arr2 []int = make([]int,l)
	var j int = 0
	for i := 0;i < l && j < l;i++{
		if arr[i] != 0{
			arr2[j] = arr[i]
			j++
		}else{
			arr2[j] = 0
			if (j + 1) < l{
				arr2[j + 1] = 0
			}
			j += 2
		}
	}
	copy(arr,arr2)
}