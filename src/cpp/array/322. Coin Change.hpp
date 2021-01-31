#include <vector>
using namespace std;

class Solution_322 {
public:
    int coinChange(vector<int>& coins, int amount) {
        int coin_count = coins.size();
        //dp[i] = min coin number when sum = i,
        std::vector<int> dp(amount + 1);
        dp[0] = 0;
        for(int i = 1;i <= amount;i++){
            dp[i] = 2147483647;
            for(int j = 0;j < coin_count;j++){
                if(i - coins[j] >= 0 && dp[i - coins[j]] != 2147483647){
                    if(dp[i - coins[j]] != 2147483647){
                        dp[i] = min(dp[i],dp[i - coins[j]] + 1);
                    }
                }
            }
        }
        if(dp[amount] == 2147483647){
            return -1;
        }
        return dp[amount];
    }
};