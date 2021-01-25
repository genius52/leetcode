#include <vector>
using namespace std;

class Solution_52 {
public:
    bool check_conflict(std::vector<std::vector<bool>>& graph,int n,int r,int col){
        if(r == 0)
            return false;
        for(int i = r - 1;i >= 0;i--){
            if(graph[i][col])
                return true;
        }
        int i = r - 1;
        int j = col - 1;
        while(i >= 0 && j >= 0){
            if(graph[i][j])
                return true;
            i--;
            j--;
        }
        i = r - 1;
        j = col + 1;
        while(i >= 0 && j < n){
            if(graph[i][j])
                return true;
            i--;
            j++;
        }
        return false;
    }

    int dfs_totalNQueens(std::vector<std::vector<bool>>& graph,int n,int row){
        if(row >= n){
            return 1;
        }
        int res = 0;
        for(int i = 0;i < n;i++){
            if(check_conflict(graph,n,row,i))
                continue;
            graph[row][i] = true;
            res += dfs_totalNQueens(graph,n,row + 1);
            graph[row][i] = false;
        }
        return res;
    }

    int totalNQueens(int n) {
        std::vector<std::vector<bool>> graph(n,std::vector<bool>(n));
        return dfs_totalNQueens(graph,n,0);
    }
};