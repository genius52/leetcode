#include <vector>

//Given an integer array with all positive numbers and no duplicates,
// find the number of possible combinations that add up to a positive integer target.
//nums = [1, 2, 3]
//target = 4

class Solution_377 {
    std::vector<int> dp;
public:
    int combinationSum4(vector<int>& nums, int target) {
        if(target == 0)
            return 1;
        if(target < 0)
            return 0;

        if(dp.empty()){
            dp.assign(target + 1,-1);
        }
        if(dp[target] != -1)
            return dp[target];
        int total = 0;
        int l = nums.size();
        for(int i = 0;i < l;i++){
            total += combinationSum4(nums,target - nums[i]);
        }
        dp[target] = total;
        return total;
    }
};