#include <vector>
#include <math.h>
using namespace std;

class Solution_598 {
public:
    int maxCount(int m, int n, vector<vector<int>>& ops) {
        int len = ops.size();
        if(len == 0)
            return m * n;
        int row = m;
        int col = n;
        for(int i = 0;i < len;i++){
            row = min(row,ops[i][0]);
            col = min(col,ops[i][1]);
        }
        return row * col;
    }
};