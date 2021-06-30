package number

//Input: k = 2, n = 6
//Output: 3

//Input: k = 3, n = 14
//Output: 4

//if(n==0||n==1) return n;  //if no. of floor 0 , 1 return n:
//        if(k==1) return n;        // if 1 egg return number of floor
//        if(memo[k][n]!=-1) return memo[k][n];
//        int ans=1000000,l=1,h=n,temp=0;
//
//        while(l<=h)
//        {
//            int mid=(l+h)/2;
//            int left=find(k-1,mid-1,memo);   //egg broken check for down floors of mid
//            int right=find(k,n-mid,memo) ;   // not broken check for up floors of mid
//            temp=1+max(left,right);          //store max of both
//            if(left<right){                  //since right is more than left and we need more in worst case
//              l=mid+1;                       // so l=mid+1 to gain more temp for worst case : upward
//            }
//            else                             //left > right so we will go downward
//            {
//                h=mid-1;
//            }
//            ans=min(ans,temp);
func dp_superEggDrop(K int,N int,memo map[int]map[int]int)int{
	if N == 0 || N == 1{
		return N
	}
	//只剩下一个蛋了，每一层都要试
	if K == 1{
		return N
	}
	if _,ok1 := memo[K];ok1 {
		if _, ok2 := memo[K][N]; ok2 {
			return memo[K][N]
		}
	}else{
		memo[K] = make(map[int]int)
	}
	var res int = 1000000
	var left int = 1
	var right int = N
	for left <= right{
		mid := (left + right)/2
		//碎了.蛋少了一个(K - 1),去下一层(i - 1)接着尝试
		res1 := dp_superEggDrop(K - 1,mid - 1,memo)
		//没碎
		res2 := dp_superEggDrop(K,N - mid,memo)
		cnt := 1 + max_int(res1,res2)
		if res1 < res2{
			left = mid + 1
		}else{
			right = mid - 1
		}
		if cnt < res{
			res = cnt
		}
	}
	memo[K][N] = res
	return res
}

func SuperEggDrop(K int, N int) int {
	var memo map[int]map[int]int = make(map[int]map[int]int)
	return dp_superEggDrop(K,N,memo)
}