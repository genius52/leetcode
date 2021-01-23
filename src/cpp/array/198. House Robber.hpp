#include <vector>
using namespace std;


class Solution_198 {
    int len = 0;
public:
    int rob(std::vector<int>& nums) {
        len = nums.size();

        if(len == 0)
            return 0;
        if(len == 1)
            return nums[0];
        int max_choose_cur = recursive(nums,2) + nums[0];
        int max_skip_cur = recursive(nums,1);
        return max_choose_cur > max_skip_cur?max_choose_cur:max_skip_cur;
    }

    int recursive(std::vector<int>& nums,int begin){
        if(begin >= len)
            return 0;
        int max_choose_cur = recursive(nums,begin + 2) + nums[begin];
        int max_skip_cur = recursive(nums,begin + 1);
        auto max = max_choose_cur > max_skip_cur?max_choose_cur:max_skip_cur;
        return max;
    }
};
