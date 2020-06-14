#include <vector>
using namespace std;
//Input: nums = [3,4,5,2]
//Output: 12
class Solution_1464 {
public:
    int maxProduct(vector<int>& nums) {
        int res = -2147483648;
        int len = nums.size();
        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (i == j)
                    continue;
                int product = (nums[i] - 1) * (nums[j] - 1);
                if (product > res)
                    res = product;
            }
        }
        return res;
    }
};