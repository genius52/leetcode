#include <vector>
#include <set>
using namespace std;

class Solution_363 {
public:
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int rows = matrix.size();
        int columns = matrix[0].size();
        std::vector<std::vector<int>> dp(rows + 1,std::vector<int>(columns + 1));
        for(int i = 1;i <= rows;i++){
            for(int j = 1;j <= columns;j++){
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1] + matrix[i - 1][j - 1];
            }
        }
        int res = -2147483648;
        for(int i = 1;i <= rows;i++){
            for(int j = i;j <= rows;j++){
                std::set<int> record;
                record.insert(0);
                for(int c = 1;c <= columns;c++){
                    int right = dp[j][c] - dp[i - 1][c];
                    auto find = record.lower_bound(right - k);
                    if(find != record.end()){
                        int cur = right - *find;
                        if(cur > res)
                            res = cur;
                    }
                    record.insert(right);
                }
            }
        }
        return res;
    }

    int maxSumSubmatrix2(vector<vector<int>>& matrix, int k) {
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
            if(dp[i][0] <= k && (k - dp[i][0]) < diff){
                diff = k - dp[i][0];
                res = dp[i][0];
            }
        }
        for(int j = 1;j < columns;j++){
            dp[0][j] = matrix[0][j] + dp[0][j - 1];
            if(dp[0][j] == k)
                return k;
            if(dp[0][j] <= k && (k - dp[0][j]) < diff){
                diff = k - dp[0][j];
                res = dp[0][j];
            }
        }
        for(int i = 1;i < rows;i++){
            for(int j = 1;j < columns;j++){
                dp[i][j] = matrix[i][j] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1];
                if(dp[i][j] == k){
                    return k;
                }
                if(dp[i][j] <= k && (k - dp[i][j]) < diff){
                    diff = k - dp[i][j];
                    res = dp[i][j];
                }
            }
        }
        std::vector<std::set<int>> record(columns);
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++) {
                int cur_area = dp[i][j];
                int diff = cur_area - k;
                for(int m = 0;m < j;m++){
                    auto find = record[m].lower_bound(diff);
                    //find--;
                    if(find != record[m].end()){
                        int cur_val = dp[i][j] - *find;
                        if(cur_val > res){
                            res = cur_val;
                        }
                    }
                }
                record[j].insert(cur_area);
            }
        }
        return res;
    }
};