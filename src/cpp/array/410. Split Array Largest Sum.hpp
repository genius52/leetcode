#include <vector>
#include <algorithm>
using namespace std;

//Given an array nums which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays.
//
//Write an algorithm to minimize the largest sum among these m subarrays.
//
//
//
//Example 1:
//
//Input: nums = [7,2,5,10,8], m = 2
//Output: 18
//Explanation:
//There are four ways to split nums into two subarrays.
//The best way is to split it into [7,2,5] and [10,8],
//where the largest sum among the two subarrays is only 18.
class Solution_410 {
public:
    int dp_splitArray(std::vector<int>& nums,int len,int pos,int m,std::unordered_map<int,std::unordered_map<int,int>>& memo){
        if(m == 1){
            return std::accumulate(nums.begin() + pos,nums.end(),0);
        }
        if(memo.find(pos) != memo.end()){
            if(memo[pos].find(m) != memo[pos].end()){
                return memo[pos][m];
            }
        }
        int min_len = 1;
        int max_len = len - pos - m + 1;
        int min_sum = 2147483647;
        for(int l = min_len;l <= max_len;l++){
            int cur_sum = std::accumulate(nums.begin() + pos,nums.begin() + pos + l,0);
            int rest_sum = dp_splitArray(nums,len,pos + l,m - 1,memo);
            int cur = max(cur_sum,rest_sum);
            if (cur < min_sum){
                min_sum = cur;
            }
        }
        memo[pos][m] = min_sum;
        return min_sum;
    }

    int splitArray(vector<int>& nums, int m) {
        int len = nums.size();
        std::unordered_map<int,std::unordered_map<int,int>> memo;
        return dp_splitArray(nums,len,0,m,memo);
    }
};