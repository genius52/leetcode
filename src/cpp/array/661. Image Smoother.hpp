#include <vector>
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
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                int cnt = 1;
                int val = M[i][j];
                if(i - 1 >= 0){
                    val += M[i - 1][j];
                    cnt++;
                    if(j - 1 >= 0){
                        val += M[i - 1][j - 1];
                        cnt++;
                    }
                    if(j + 1 < columns){
                        val += M[i - 1][j + 1];
                        cnt++;
                    }
                }
                if(i + 1 < rows){
                    val += M[i + 1][j];
                    cnt++;
                    if(j - 1 >= 0){
                        val += M[i + 1][j - 1];
                        cnt++;
                    }
                    if(j + 1 < columns){
                        val += M[i + 1][j + 1];
                        cnt++;
                    }
                }
                if(j - 1 >= 0){
                    val += M[i][j - 1];
                    cnt++;
                }
                if(j + 1 < columns){
                    val += M[i][j + 1];
                    cnt++;
                }
                res[i][j] = val / cnt;
            }
        }
        return res;
    }
};