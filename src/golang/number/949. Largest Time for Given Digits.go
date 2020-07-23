package number

import (
	"fmt"
)

//Input: [1,2,3,4]
//Output: "23:41"
func perm_largestTimeFromDigits(A []int,pos int,target int) (bool,int){
	if(pos == target){
		if((A[0] == 0 || A[0] == 1) && A[2] <= 5 && A[3] <= 9){
			return true,1000 * A[0] + 100 * A[1] + 10 * A[2] + A[3];
		}
		if(A[0] == 2 && A[1] <=3 && A[2] <= 5 && A[3] <= 9){
			return true,1000 * A[0] + 100 * A[1] + 10 * A[2] + A[3];
		}
		return false,0
	}else{
		var max int = -1
		var ret bool = false
		for i := pos;i < len(A);i++{
			A[i],A[pos] = A[pos],A[i]
			r,val := perm_largestTimeFromDigits(A,pos + 1,target)
			if(r && val > max){
				ret = r
				max = val
			}
			A[i],A[pos] = A[pos],A[i]
		}
		return ret,max
	}
}

func LargestTimeFromDigits(A []int) string {
	ret,val := perm_largestTimeFromDigits(A,0,3)
	if(ret){
		 //strconv.Itoa(val/100) + ":" + strconv.Itoa(val%100)
		 res := fmt.Sprintf("%.2d:%.2d",val/100,val%100)
		 return res
	}
	return ""
}
