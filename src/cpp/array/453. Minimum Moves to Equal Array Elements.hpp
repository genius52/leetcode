#include <vector>
#include <algorithm>
using namespace std;

class Solution_453 {
public:
    int minMoves(vector<int>& nums) {
        std::sort(nums.begin(),nums.end());
        int res = 0;
        int begin = 0;
        int end = nums.size() - 1;
        while(begin < end && nums[begin] != nums[end]){
            res += (nums[end] - nums[begin]);
            end--;
        }
        return res;
    }
};