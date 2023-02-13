#include <vector>
#include <set>
using namespace std;

class Solution_2563 {
public:
    long long countFairPairs(vector<int>& nums, int lower, int upper) {
        std::sort(nums.begin(),nums.end());
        int len = nums.size();
        long long res = 0;
        for(int i = 0;i < len;i++){
            int find_low = lower - nums[i];
            int find_high = upper - nums[i];
            auto it1 = std::lower_bound(nums.begin() + i + 1,nums.end(),find_low);
            if(it1 == nums.end()){
                continue;
            }
            auto it2 = std::upper_bound(nums.begin() + i + 1,nums.end(),find_high);
            int dis1 = std::distance(nums.begin(),it1);
            int dis2 = std::distance(nums.begin(),it2);
            if(dis1 <= dis2)
                res += std::distance(it1, it2);
        }
        return res;
    }
};