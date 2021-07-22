package array


//969
//Input: [3,2,4,1]
//Output: [4,2,4,3]
//Explanation:
//We perform 4 pancake flips, with k values 4, 2, 4, and 3.
//Starting state: A = [3, 2, 4, 1]
//After 1st flip (k=4): A = [1, 4, 2, 3]
//After 2nd flip (k=2): A = [4, 1, 2, 3]
//After 3rd flip (k=4): A = [3, 2, 1, 4]
//After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
func reverse_pancakeSort(A []int,begin int,end int){
	for begin < end{
		A[begin],A[end] = A[end],A[begin]
		begin++
		end--
	}
}

func PancakeSort2(A []int) []int{
	var l int = len(A)
	var res []int
	for i := l - 1;i > 0;i--{
		for j := i;j >= 0;j--{
			if A[j] == (i + 1){
				reverse_pancakeSort(A,0,j)
				res = append(res,j + 1)
				reverse_pancakeSort(A,0,i)
				res = append(res,i + 1)
				break
			}
		}
	}
	return res
}

func pancakeSort(A []int) []int {
	l := len(A)
	var res []int
	for i := l - 1;i >= 0;i--{
		for j := 0;j <= i;j++{
			if A[j] == i + 1{
				if j != 0{
					res = append(res,j + 1)
					reverse_pancakeSort(A,0,j)
				}
				if i != 0{
					reverse_pancakeSort(A,0,i)
					res = append(res,i + 1)
				}
				break
			}
		}
	}
	return res
}