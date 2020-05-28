#include <vector>
#include <set>
using namespace std;

//Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements in this subset satisfies:
//
//Si % Sj = 0 or Sj % Si = 0.

//Input: [1,2,4,8]
//Output: [1,2,4,8]
class Solution_368 {
public:
    vector<int> largestDivisibleSubset(vector<int>& nums) {
        std::vector<int> res;
        std::sort(nums.begin(),nums.end());
        int len = nums.size();
        if(len <= 1)
            return nums;
        if(len == 2){
            if(nums[1] % nums[0] != 0){
                return {nums[0]};
            }
        }
        int max_len = 0;
        for(int i = len - 1;i > 0 ;i--) {
            for(int j = i - 1;j >= 0;j--){
                if(nums[i] % nums[j] != 0){
                    continue;
                }
                int last_num = nums[j];
                std::vector<int> v{nums[i],nums[j]};
                for(int k = j - 1;k >= 0;k--){
                    if(last_num % nums[k] == 0){
                        v.push_back(nums[k]);
                        last_num = nums[k];
                    }
                }
                if(v.size() > max_len){
                    res = v;
                    max_len = v.size();
                }
                v.clear();
            }
        }
        if(max_len == 0){
            return {nums[0]};
        }
        return res;
    }
};