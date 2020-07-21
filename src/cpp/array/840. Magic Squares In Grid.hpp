#include <unordered_set>
using namespace std;
class Solution_840 {
public:
    bool isvalid(vector<vector<int>>& grid,int row,int column){
        int sum = grid[row - 1][column - 1] + grid[row][column] + grid[row + 1][column + 1];
        if((grid[row + 1][column - 1] + grid[row][column] + grid[row - 1][column + 1]) != sum)
            return false;
        for(int i = row - 1;i <= row + 1;i++){
            if(sum != (grid[i][column - 1] + grid[i][column] + grid[i][column + 1]))
                return false;
        }
        for(int j = column - 1;j < column + 1;j++){
            if(sum != (grid[row - 1][j] + grid[row][j] + grid[row + 1][j]))
                return false;
        }
        std::unordered_set<int> s;
        s.insert(grid[row - 1][column - 1]);
        s.insert(grid[row - 1][column]);
        s.insert(grid[row - 1][column + 1]);
        s.insert(grid[row][column - 1]);
        s.insert(grid[row][column]);
        s.insert(grid[row][column + 1]);
        s.insert(grid[row + 1][column - 1]);
        s.insert(grid[row + 1][column]);
        s.insert(grid[row + 1][column + 1]);
        for(int i = 1;i <= 9;i++){
            if(s.find(i) == s.end())
                return false;
        }
        return true;
    }

    int numMagicSquaresInside(vector<vector<int>>& grid) {
        int res = 0;
        int rows = grid.size();
        int columns = grid[0].size();
        for(int i = 1;i < rows - 1;i++){
            for(int j = 1;j < columns - 1;j++){
                if(isvalid(grid,i,j))
                    res++;
            }
        }
        return res;
    }
};