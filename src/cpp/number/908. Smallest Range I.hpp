#include <math.h>

class Solution_908 {
public:
    int smallestRangeI(vector<int>& A, int K) {
        int small = INT_MAX;
        int big = INT_MIN;
        for(auto n: A){
            if(n < small)
                small = n;
            if(n > big)
                big = n;
        }
        if(abs(big - small) <= K * 2)
            return 0;
        int diff = abs(big - small) - K * 2;
        return diff;
    }
};