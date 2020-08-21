#include <vector>
#include <string>
#include <math.h>
using namespace std;


//Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
//Output: 4
//Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s, which are "10","0001","1","0".
class Solution_474 {
public:
    int dp_findMaxForm(vector<string>& strs,int pos,int zero_left,int one_left, std::vector<std::vector<std::vector<int>>>& memo){
        int len = strs.size();
        if(pos >= len || (zero_left <= 0 && one_left <= 0)){
            return 0;
        }
        if(memo[zero_left][one_left][pos] != 0){
            return memo[zero_left][one_left][pos];
        }
        int zero_cnt = 0;
        int one_cnt = 0;
        for(auto c: strs[pos]){
            if(c == '0')
                zero_cnt++;
            else
                one_cnt++;
        }
        int size = 0;
        if(zero_cnt <= zero_left && one_cnt <= one_left){
            size = 1 + dp_findMaxForm(strs,pos + 1,zero_left - zero_cnt,one_left - one_cnt,memo);
        }
        auto res = max(size,dp_findMaxForm(strs,pos + 1,zero_left,one_left,memo));
        memo[zero_left][one_left][pos] = res;
        return res;
    }

    int findMaxForm(vector<string>& strs, int m, int n) {
//        std::vector<std::vector<int>> dp;//dp[i][j] = max size numbers with i '0's and j '1's
//        dp.resize(m + 1);
//        for(int i = 0;i <= m;i++){
//            dp[i].resize(n + 1);
//        }
//        int len = strs.size();
//        for(auto s : strs){
//            int zero_cnt = 0;
//            int one_cnt = 0;
//            for(auto c : s){
//                if(c == '0')
//                    zero_cnt++;
//                else
//                    one_cnt++;
//            }
//            if(zero_cnt <= m && one_cnt <= n){
//                for(int i = m;i >= zero_cnt;i--){
//                    for(int j = n;j >= one_cnt;j--){
//                        dp[i][j] = max(dp[i][j],1 + dp[i - zero_cnt][j - one_cnt]);
//                    }
//                }
//            }
//        }
//        return dp[m][n];

        int len = strs.size();
        std::vector<std::vector<std::vector<int>>> graph;
        graph.resize(m + 1);
        for(int i = 0;i <= m;i++){
            graph[i].resize(n + 1);
        }
        for(int i = 0;i <= m;i++){
            for(int j = 0;j <= n;j++){
                graph[i][j].resize(len + 1);
            }
        }
//        std::unordered_map<int,int> memo;
//        std::sort(strs.begin(), strs.end(), [](string &a, string &b){
//            return a.size() < b.size();
//        });
        return dp_findMaxForm(strs,0,m,n,graph);
    }
};