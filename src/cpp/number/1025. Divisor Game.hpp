#include <vector>
using namespace std;
//Input: 3
//Output: false
//Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
class Solution_1025 {
public:
    bool divisorGame(int N) {
        return (N | 1) != N;
    }
    bool divisorGame2(int N) {
        std::vector<bool> dp(N + 1);
        dp[1] = false;
        dp[2] = true;
        for(int i = 3;i <= N;i++){
            for(int j = 1;j < i;j++){
                if(i % j == 0 && !dp[i - j]){
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[N];
    }
};