#include <vector>
using namespace std;

class Solution_1770 {
public:
    int dp_maximumScore(vector<int>& nums, int begin, int end, vector<int>& multipliers, int pos, int multipliers_len, std::vector<std::vector<int>>& memo) {
        if (begin + end >= multipliers_len) {
            return 0;
        }
        if (memo[begin][end] != 0) {
            return memo[begin][end];
        }
        if (pos >= multipliers_len)
            return 0;
        //choose begin
        int choose_begin = nums[begin] * multipliers[pos] + dp_maximumScore(nums, begin + 1, end, multipliers, pos + 1, multipliers_len, memo);
        //choose end
        int len = nums.size();
        int choose_end = nums[len - 1 - end] * multipliers[pos] + dp_maximumScore(nums, begin, end + 1, multipliers, pos + 1, multipliers_len, memo);
        memo[begin][end] = max(choose_begin, choose_end);
        return memo[begin][end];
    }

    int maximumScore(vector<int>& nums, vector<int>& multipliers) {
        int nums_len = nums.size();
        int multipliers_len = multipliers.size();
        std::vector<std::vector<int>> memo(1001, std::vector<int>(1001));
        return dp_maximumScore(nums, 0, 0, multipliers, 0, multipliers_len, memo);
    }
};