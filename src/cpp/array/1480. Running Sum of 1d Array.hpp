#include <vector>
using namespace std;

class Solution_1480 {
public:
    vector<int> runningSum(vector<int>& nums) {
        int len = nums.size();
        std::vector<int> res;
        res.resize(len);
        res[0] = nums[0];
        for (int i = 1; i < len; i++) {
            res[i] = res[i - 1] + nums[i];
        }
        return res;
    }
};