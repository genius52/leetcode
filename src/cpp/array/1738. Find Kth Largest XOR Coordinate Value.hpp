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
        std::vector<int> data(rows * columns);
        data[0] = dp[0][0];
        int index = 1;
        for(int i = 1;i < rows;i++){
            dp[i][0] = dp[i - 1][0] ^ matrix[i][0];
            data[index++] = dp[i][0];
        }
        for(int j = 1;j < columns;j++){
            dp[0][j] = dp[0][j - 1] ^ matrix[0][j];
            data[index++] = dp[0][j];
        }
        for(int i = 1;i < rows;i++){
            for(int j = 1;j < columns;j++){
                dp[i][j] = matrix[i][j] ^ dp[i - 1][j] ^ dp[i][j - 1] ^ dp[i - 1][j - 1];
                data[index++] = dp[i][j];
            }
        }
        std::partial_sort(data.begin(), data.begin() + k,data.end() , std::greater<>());
        return data[k - 1];
    }
};