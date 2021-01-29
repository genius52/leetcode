//Input: 2
//Output: 91
//Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100,
//             excluding 11,22,33,44,55,66,77,88,99
#include <vector>
using namespace std;

class Solution_357 {
public:
    int countNumbersWithUniqueDigits(int n) {
        if (n == 0){
            return 1;
        }
        std::vector<int> dp(n + 1);
        dp[0] = 1;
        dp[1] = 10;
        for (int i = 2;i <= n;i++){
            dp[i] += dp[i - 1];
            int k = 0;
            int cnt = 9;
            for(int j = i - 1;j > 0;j--){
                cnt *= (9 - k);
                k++;
            }
            dp[i] += cnt;
        }
        return dp[n];
    }
};