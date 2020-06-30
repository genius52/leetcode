#include <vector>
#include <set>
#include <unordered_set>

using namespace std;


class Solution_1498 {
public:
    int numSubseq(vector<int>& nums, int target) {
        int len = nums.size();
        if(len == 0)
            return 0;
        int mod = (int)1e9 + 7;
        std::sort(nums.begin(),nums.end());
        int total = 0;
        std::vector<int> record(len,1);
        for(int i = 1;i < len;i++){
            record[i] = record[i - 1] * 2 % mod;
        }
        int l = 0;
        int r = len - 1;
        while(l < len){
            if(nums[l] * 2 > target)
                break;
            while(r > l && nums[l] + nums[r] > target){
                r--;
            }
            total = (total + record[r - l]) % mod;
            l++;
        }
        return total;
    }
};