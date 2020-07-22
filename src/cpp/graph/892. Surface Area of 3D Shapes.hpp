#include <vector>
#include <math.h>
using namespace std;

//Input: [[1,2],[3,4]]
//Output: 34
class Solution_892 {
public:
    int surfaceArea(vector<vector<int>>& grid) {
        int rows = grid.size();
        int columns = grid[0].size();
        int res = 0;
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                if(grid[i][j] == 0)
                    continue;
                res += 2 + grid[i][j] * 4;
                int dups = 0;
                if(i - 1 >= 0){
                    dups += min(grid[i - 1][j],grid[i][j]);
                }
                if(j - 1 >= 0){
                    dups += min(grid[i][j - 1],grid[i][j]);
                }
                res = res - 2 * dups;
            }
        }
        return res;
    }
};