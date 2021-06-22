#include <vector>
#include <map>
using namespace std;

//Input: nums = [2,-1,2], k = 3
//Output: 3
class Solution_862 {
public:
    int shortestSubarray(vector<int>& nums, int k) {
        int len = nums.size();
        std::map<int,int,std::less<int>> sum_index;
        sum_index[0] = -1;
        int res = 2147483647;
        int sum = 0;
        for(int i = 0;i < len;i++){
            sum += nums[i];
            int target = sum - k;
            auto find = sum_index.upper_bound(target);//小于等于target
            for(auto it = sum_index.begin();it != find;){
                int cur_range = i - (*it).second;
                if(cur_range < res){
                    res = cur_range;
                }
                sum_index.erase(it++);
            }
            sum_index[sum] = i;
        }
        if(res == 2147483647)
            return -1;
        return res;
    }
};