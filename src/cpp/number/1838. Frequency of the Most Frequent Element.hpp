#include <vector>
#include <map>
using namespace std;

class Solution_1838 {
public:
    int maxFrequency(vector<int>& nums, int k) {
        int len = nums.size();
        std::sort(nums.begin(),nums.end());
        int max_len = 0;
        int left = 0;
        while(left < len){
            int right = left + 1;
            int pre_num = nums[left];
            int diff = 0;
            while(right < len){
                if(nums[right] == nums[right - 1]){
                    right++;
                    continue;
                }
                int cnt = right - left;
                diff += (nums[right] - pre_num) * cnt;
                pre_num = nums[right];
                if(diff > k)
                    break;
                right++;
            }
            if ((right - left) > max_len){
                max_len = right - left;
            }
            left++;
        }
        return max_len;
    }
};