#include <vector>
#include <queue>
using namespace std;

class Solution_1696 {
public:
    int maxResult(vector<int>& nums, int k) {
        int len = nums.size();
        //std::vector<int> dp(len);//dp[i]跳到i的最大和
        std::priority_queue<std::pair<int,int>> q;//最大和,索引
        //dp[0] = nums[0];
        int max_val = nums[0];
        q.push({nums[0],0});
        for (int i = 1; i < len; ++i) {
            while(!q.empty() && i - q.top().second > k){
                q.pop();
            }
            max_val = nums[i] + q.top().first;
            q.push({max_val,i});
        }
        return max_val;
    }
};