#include <vector>
#include <set>
using namespace std;

class Solution {
public:
    int longestSubarray(vector<int>& nums, int limit) {
        int len = nums.size();
        if (len <= 1)
            return len;
        int max_len = 1;
        std::multiset<int> set;
        int left = 0;
        for (int i = 0; i < len; i++) {
            set.insert(nums[i]);
            while (abs(*set.begin() - *set.rbegin()) > limit) {
                set.erase(nums[left]);
                left++;
            }
            if (set.size() > max_len)
                max_len = set.size();
        }
        return max_len;
    }
};