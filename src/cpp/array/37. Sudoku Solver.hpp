#include <vector>
#include <unordered_map>
using namespace std;

//Input: board =
//[["5","3",".",".","7",".",".",".","."],
// ["6",".",".","1","9","5",".",".","."],
// [".","9","8",".",".",".",".","6","."],
// ["8",".",".",".","6",".",".",".","3"],
// ["4",".",".","8",".","3",".",".","1"],
// ["7",".",".",".","2",".",".",".","6"],
// [".","6",".",".",".",".","2","8","."],
// [".",".",".","4","1","9",".",".","5"],
// [".",".",".",".","8",".",".","7","9"]]
//Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
class Solution_37 {
public:
    bool is_valid(vector<vector<char>>& board,int row,int col,int val){
        for(int i = 0;i < 9;i++){//check same columns
            if(i == row)
                continue;
            if(board[i][col] == val)
                return false;
        }
        for(int j = 0;j < 9;j++){//check same row
            if(j == col)
                continue;
            if(board[row][j] == val)
                return false;
        }
//        if(row == col){//check tilt
//            for(int i = 0;i < 9;i++){
//                if(i == row)
//                    continue;
//                if(board[i][i] == val)
//                    return false;
//            }
//        }
        int row_palace = row/3 * 3;
        int col_palace = col/3 * 3;
        for(int i = row_palace;i < row_palace + 3;i++){
            for(int j = col_palace;j < col_palace + 3;j++){
                if(i == row && j == col)
                    continue;
                if(board[i][j] == val)
                    return false;
            }
        }
        return true;
    }

    bool dfs_solveSudoku(vector<vector<char>>& board,std::vector<std::vector<char>>& origin,int pos){
        int row = pos / 9;
        int col = pos % 9;
        if(row == 9)
            return true;
        if(origin[row][col] != '.')
            return dfs_solveSudoku(board,origin,pos + 1);
        for(int i = 1;i <= 9;i++){
            char n = i + '0';
            if(is_valid(board,row,col,n)){
                board[row][col] = n;
                auto ret = dfs_solveSudoku(board,origin,pos + 1);
                if(ret){
                    return true;
                }else{
                    board[row][col] = '.';
                }
            }
        }
        return false;
    }

    void solveSudoku(vector<vector<char>>& board) {
        std::vector<std::vector<char>> origin = board;
        dfs_solveSudoku(board,origin,0);
    }
};