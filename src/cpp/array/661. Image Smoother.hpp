#include <vector>
#include <array>
using namespace std;

class Solution_661 {
public:
    vector<vector<int>> imageSmoother(vector<vector<int>>& M) {
        int rows = M.size();
        int columns = M[0].size();
        std::vector<std::vector<int>> res;
        res.resize(rows);
        for(int i = 0;i < rows;i++){
            res[i].resize(columns);
        }
        std::array<std::array<int,2>,8> dir{{{0,-1},{-1,-1},{-1,0},{-1,1},
                                                    {0,1},{1,1},{1,0},{1,-1}}};
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                int cnt = 1;
                int val = M[i][j];
                for(auto d : dir){
                    int x = i + d[0];
                    int y = j + d[1];
                    if(x >= 0 && x < rows && y >= 0 && y < columns){
                        val += M[x][y];
                        cnt++;
                    }
                }
                res[i][j] = val / cnt;
            }
        }
        return res;
    }
};