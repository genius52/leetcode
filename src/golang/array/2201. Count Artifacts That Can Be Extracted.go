package array

func DigArtifacts(n int, artifacts [][]int, dig [][]int) int {
	var arr [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		arr[i] = make([]int,n)
	}
	var cnt int = 1
	for _,pos := range artifacts{
		for i := pos[0];i <= pos[2];i++{
			for j := pos[1];j <= pos[3];j++{
				arr[i][j] = cnt
			}
		}
		cnt++
	}
	for _,d := range dig{
		arr[d[0]][d[1]] = 0
	}
	var record []bool = make([]bool,cnt)
	for i := 0;i < n;i++{
		for j := 0;j < n;j++{
			if arr[i][j] != 0{
				record[arr[i][j]] = true
			}
		}
	}
	var res int = 0
	for i := 1;i < cnt;i++{
		if !record[i]{
			res++
		}
	}
	return res
}