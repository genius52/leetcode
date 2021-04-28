#include <vector>
using namespace std;

class Solution_628 {
public:
    int maximumProduct(vector<int>& nums) {
        int len = nums.size();
        std::sort(nums.begin(),nums.end());
        return max(nums[0] * nums[1] * nums[len - 1],nums[len - 1] * nums[len - 2] * nums[len - 3]);
    }
};