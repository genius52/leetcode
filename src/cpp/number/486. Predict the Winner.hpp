#include <vector>
using namespace std;

//Input: [1, 5, 233, 7]
//Output: True
//Explanation: Player 1 first chooses 1. Then player 2 have to choose between 5 and 7. No matter which number player 2 choose, player 1 can choose 233.
//Finally, player 1 has more score (234) than player 2 (12), so you need to return True representing player1 can win.
class Solution_486 {
public:
    int dp_PredictTheWinner(vector<int>& nums,int left,int right,std::vector<std::vector<int>>& memo){
        if(left > right)
            return 0;
        if(left == right)
            return nums[left];
        if(memo[left][right] != 0){
            return memo[left][right];
        }
        auto choose_left = nums[left] + min(dp_PredictTheWinner(nums,left + 2,right,memo),dp_PredictTheWinner(nums,left + 1,right - 1,memo));
        auto choose_right = nums[right] +  min(dp_PredictTheWinner(nums,left + 1,right - 1,memo),dp_PredictTheWinner(nums,left,right - 2,memo));
        memo[left][right] = max(choose_left,choose_right);
        return memo[left][right];
    }

    bool PredictTheWinner(vector<int>& nums) {
        int len = nums.size();
        int total = 0;
        for(auto n : nums){
            total += n;
        }
        std::vector<std::vector<int>> memo(len,std::vector<int>(len));
        auto score1 = dp_PredictTheWinner(nums,0,len - 1,memo);
        return score1 >= (total - score1);
    }
};