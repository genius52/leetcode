#include <vector>
#include <set>
using namespace std;

class Solution_363 {
public:
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int rows = matrix.size();
        int columns = matrix[0].size();
        std::vector<std::vector<int>> dp(rows,std::vector<int>(columns));
        dp[0][0] = matrix[0][0];
        int res = 0;
        int diff = 2147483647;
        for(int i = 1;i < rows;i++){
            dp[i][0] = matrix[i][0] + dp[i - 1][0];
            if(dp[i][0] == k)
                return k;
        }
        for(int j = 1;j < columns;j++){
            dp[0][j] = matrix[0][j] + dp[0][j - 1];
            if(dp[0][j] == k)
                return k;
        }
        for(int i = 1;i < rows;i++){
            for(int j = 1;j < columns;j++){
                dp[i][j] = matrix[i][j] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1];
                if(dp[i][j] == k){
                    return k;
                }
            }
        }

        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                for(int r = 0;r <= i;r++){
                    for(int c = 0;c <= j;c++){
                        int cur_area = dp[i][j];
                        if(c != 0){
                            cur_area -= dp[i][c - 1];
                        }
                        if(r != 0){
                            cur_area -= dp[r - 1][j];
                        }
                        if(c != 0 && r != 0){
                            cur_area += dp[r - 1][c - 1];
                        }
                        if(cur_area == k)
                            return k;
                        if(cur_area <= k && (k - cur_area) < diff){
                            diff = k - cur_area;
                            res = cur_area;
                        }
                    }
                }
            }
        }
        return res;
    }
};