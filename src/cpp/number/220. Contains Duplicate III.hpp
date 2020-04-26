#include <vector>
#include <map>
#include <math.h>

//Given an array of integers, find out whether there are two distinct indices i and j in the array 
//such that the absolute difference between nums[i] and nums[j] is at most t and 
//the absolute difference between i and j is at most k.
//Input: nums = [1,2,3,1], k = 3, t = 0
//Output: true
class Solution_220 {
public:
    bool containsNearbyAlmostDuplicate(std::vector<int>& nums, int k, int t) {
        if (t < 0){
            return false;
        }
        std::map<long, long> record;
        for (int i = 0; i < nums.size(); i++) {
            if (record.empty()) {
                record[nums[i]] = i;
                continue;
            }
            long upper = long(nums[i]) + long(t);
            long lower = long(nums[i]) - long(t);
            auto find1 = record.lower_bound(lower);
            auto find2 = record.upper_bound(upper);
            auto visit = find1;
            while (visit != find2){
                if (i - visit->second <= k)
                    return true;
                visit++;
            }
            record[nums[i]] = i;
        }
        return false;
    }
};