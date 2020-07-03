#include <vector>
#include <algorithm>
#include <math.h>
using namespace std;

//Input:
//[1,2,3]
//
//Output:
//2
//
//Explanation:
//Only two moves are needed (remember each move increments or decrements one element):
//
//[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
class Solution_462 {
public:
    int minMoves2(vector<int>& nums) {
        int len = nums.size();
        if(len <= 1)
            return 0;
        //std::sort(nums.begin(),nums.end());
        int res = 0;
        std::nth_element(nums.begin(),nums.begin() + len / 2, nums.end());
        int mid = nums[len/2];
        for(int i = 0;i < len;i++){
            res += abs(nums[i] - mid);
        }
        return res;
    }
};