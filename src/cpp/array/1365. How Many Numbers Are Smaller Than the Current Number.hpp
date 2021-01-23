#include <vector>
#include <algorithm>
using namespace std;

//Input: nums = [8,1,2,2,3]
//Output: [4,0,1,1,3]
class Solution_1365 {
public:
    std::vector<int> smallerNumbersThanCurrent(std::vector<int>& nums) {
        int len = nums.size();
        std::vector<int> res(len);
        std::vector<int> sorted_data;
        sorted_data = nums;
        std::sort(sorted_data.begin(),sorted_data.end());
        for(int i = 0;i < len;i++){
            auto it  = std::lower_bound(sorted_data.begin(),sorted_data.end(),nums[i]);
            res[i] = it - sorted_data.begin();
        }
        return res;
    }
};