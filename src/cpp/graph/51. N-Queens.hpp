#include <vector>
using namespace std;

class Solution_51 {
public:
    bool check_conflict(std::vector<std::string>& pre,int n,int r,int col){
        int len = pre.size();
        if(len == 0)
            return false;
        for(int i = r - 1;i >= 0;i--){
            if(pre[i][col] == 'Q')
                return true;
        }
        int i = r - 1;
        int j = col - 1;
        while(i >= 0 && j >= 0){
            if(pre[i][j] == 'Q')
                return true;
            i--;
            j--;
        }
        i = r - 1;
        j = col + 1;
        while(i >= 0 && j < n){
            if(pre[i][j] == 'Q')
                return true;
            i--;
            j++;
        }
        return false;
    }

    void dfs_solveNQueens(std::vector<string>& pre,int n,int r,std::vector<std::vector<std::string>>& res){
        if(r >= n){
            res.push_back(pre);
            return;
        }
        for(int col = 0;col < n;col++){
            if(check_conflict(pre,n,r,col)){
                continue;
            }
            std::string cur(n,'.');
            cur[col] = 'Q';
            std::vector<string> next = pre;
            next.push_back(cur);
            dfs_solveNQueens(next,n,r + 1,res);
        }
    }

    vector<vector<string>> solveNQueens(int n) {
        std::vector<std::vector<std::string>> res;
        std::vector<string> cur;
        dfs_solveNQueens(cur,n,0,res);
        return res;
    }
};