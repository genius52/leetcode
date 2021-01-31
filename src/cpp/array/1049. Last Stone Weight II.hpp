#include <vector>
#include <queue>
using namespace std;


class Solution_1049 {
public:
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