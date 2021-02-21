#include <vector>
#include <algorithm>
using namespace std;

class Solution_1738 {
public:
    int kthLargestValue(vector<vector<int>>& matrix, int k) {
        int rows = matrix.size();
        int columns = matrix[0].size();
        std::vector<std::vector<int>> dp(rows,std::vector<int>(columns));
        dp[0][0] = matrix[0][0];
        std::priority_queue<int,std::vector<int>,std::greater<int>> q;
        q.push(dp[0][0]);
        int index = 1;
        for(int i = 1;i < rows;i++){
            dp[i][0] = dp[i - 1][0] ^ matrix[i][0];
            q.push(dp[i][0]);
            if(q.size() > k){
                q.pop();
            }
            //data[index++] = dp[i][0];
        }
        for(int j = 1;j < columns;j++){
            dp[0][j] = dp[0][j - 1] ^ matrix[0][j];
            q.push(dp[0][j]);
            if(q.size() > k){
                q.pop();
            }
        }
        for(int i = 1;i < rows;i++){
            for(int j = 1;j < columns;j++){
                dp[i][j] = matrix[i][j] ^ dp[i - 1][j] ^ dp[i][j - 1] ^ dp[i - 1][j - 1];
                q.push(dp[i][j]);
                if(q.size() > k){
                    q.pop();
                }
            }
        }
        return q.top();
    }
};