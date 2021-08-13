#include <vector>
#include <queue>
using namespace std;


class Solution_1049 {
public:
    int lastStoneWeight(vector<int>& stones){
        int len = stones.size();
        std::priority_queue<int> q;
        for(int i = 0;i < len;i++){
            q.push(stones[i]);
        }
        while(q.size() > 1){
            auto top1 = q.top();
            q.pop();
            auto top2 = q.top();
            q.pop();
            auto diff = abs(top1 - top2);
            if(diff != 0)
                q.push(diff);
        }
        if(q.size() == 0)
            return 0;
        return q.top();
    }
    
    //1 <= stones.length <= 30
    //1 <= stones[i] <= 100
    int lastStoneWeightII(vector<int>& stones){
        int len = stones.size();
        int sum = 0;
        for(int i = 0;i < len;i++){
            sum += stones[i];
        }
        //dp[i] = sum to number i is possible
        int target = sum/2;
        std::vector<bool> dp(sum + 1);
        int cur_sum = 0;
        for(int i = 0;i < len;i++){
            cur_sum += stones[i];
            for(int j = 1;j <= cur_sum;j++){
                if(dp[j]){
                    dp[cur_sum - j] = true;
                }
            }
            dp[cur_sum] = true;
        }
        int res = 0;
        for(int i = target;i >= 0;i--){
            if(dp[i]){
                res = i;
                break;
            }
        }
        return sum - res * 2;
    }
};