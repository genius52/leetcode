package array

//Example 1:
//
//Input: A = [1,0,2]
//Output: true
//Explanation: There is 1 global inversion, and 1 local inversion.
//Example 2:
//
//Input: A = [1,2,0]
//Output: false
//Explanation: There are 2 global inversions, and 1 local inversion.
func IsIdealPermutation(A []int) bool {
	var l int = len(A)
	for i := 0;i < l;i++{
		for j := i + 2;j < l;j++{
			if A[i] > A[j]{
				return false
			}
		}
	}
	return true
}