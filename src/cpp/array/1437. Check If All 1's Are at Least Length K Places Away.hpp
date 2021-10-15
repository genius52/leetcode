#include <vector>
using namespace std;

class Solution {
public:
    bool kLengthApart(vector<int>& nums, int k) {
        if (k <= 0)
            return true;
        int len = nums.size();
        int last = 0;
        while (last < len) {
            if (nums[last] == 0)
                last++;
            else
                break;
        }
        int cur = last + 1;
        while (cur < len) {
            if (nums[cur] == 0) {
                cur++;
                continue;
            }
            if ((cur - last - 1) < k)
                return false;
            else {
                last = cur;
                cur++;
            }
        }
        return true;
    }
};