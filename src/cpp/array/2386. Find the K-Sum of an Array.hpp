#include <vector>
#include <queue>
using namespace std;

class Solution_2386 {
public:
    long long kSum(vector<int>& nums, int k) {
        int len = nums.size();
        long long max_sum = 0;
        for(int i = 0;i < len;i++){
            if(nums[i] >= 0){
                max_sum += nums[i];
            }else{
                nums[i] = -nums[i];
            }
        }
        std::sort(nums.begin(),nums.end());
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>, std::less<std::pair<long long,int>>> big_top;
        big_top.push({max_sum,0});
        while(!big_top.empty() && k > 1){
            auto [sum,idx] = big_top.top();
            big_top.pop();
            k--;
            if(idx < len){
                big_top.push({ sum - nums[idx],idx + 1});
                if(idx - 1 >= 0){
                    big_top.push({sum - nums[idx] + nums[idx - 1],idx + 1});
                }
            }
        }
        return big_top.top().first;
    }
};