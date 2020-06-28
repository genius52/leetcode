#include <math.h>

//Input: n = 12, k = 3
//Output: 3
//Explanation: Factors list is [1, 2, 3, 4, 6, 12], the 3rd factor is 3.
class Solution_1492 {
public:
    int kthFactor(int n, int k) {
        int limit = n / 2;
        int cnt = 0;
        for(int i = 1;i <= limit && cnt <= k;i++){
            if(n % i == 0){
                cnt++;
            }
            if(cnt == k)
                return i;
        }
        if(cnt == k - 1)
            return n;
        return -1;
    }
};