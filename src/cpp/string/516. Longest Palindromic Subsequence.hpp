#include <string>
#include <vector>
using namespace std;

//Input: s = "bbbab"
//Output: 4
//Explanation: One possible longest palindromic subsequence is "bbbb".
class Solution_516 {
public:
    int longestPalindromeSubseq(string s) {
        int len = s.size();
        //dp[left][right] = max length from s[left] to s[right]
        std::vector<std::vector<int>> dp(len,std::vector<int>(len));
        for(int i = 0;i < len;i++){
            dp[i][i] = 1;
        }
        for(int right = 0;right < len;right++){
            for(int left = right - 1;left >= 0;left--){
                if(s[left] != s[right]){
                    dp[left][right] = max(dp[left + 1][right],dp[left][right - 1]);
                }else{
                    if(left == right - 1){
                        dp[left][right] = 2;
                    }else{
                        dp[left][right] = 2 + dp[left + 1][right - 1];
                    }
                }
            }
        }
        return dp[0][len - 1];
    }
};