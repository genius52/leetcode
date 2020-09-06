#include <vector>
using namespace std;

class Solution_1572 {
public:
    int diagonalSum(vector<vector<int>>& mat) {
        int len = mat.size();
        int result = 0;
        for(int i = 0;i < len;i++){
            result += mat[i][i];
            result += mat[len - 1 - i][i];
        }
        if (len % 2 == 1){
            result -= mat[len/2][len/2];
        }
        return result;
    }
};