#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

//Given the following 5x5 matrix:
//
//  Pacific ~   ~   ~   ~   ~
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *
//          *   *   *   *   * Atlantic
//
//Return:
//
//[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).
class Solution_417 {
public:
    void bfs_pacificAtlantic(vector<vector<int>>& matrix,std::queue<std::pair<int,int>>& q,vector<vector<bool>>& visited){
        int rows = matrix.size();
        int columns = matrix[0].size();
        std::vector<std::pair<int,int>> dirs = {{-1,0},{1,0},{0,-1},{0,1}};
        while (!q.empty()){
            auto pos = q.front();
            q.pop();
            for(auto dir:dirs){
                int x = pos.first + dir.first;
                int y = pos.second + dir.second;
                if(x < 0 || x >= rows || y < 0 || y >= columns)
                    continue;
                if(visited[x][y])
                    continue;
                auto cur = matrix[x][y];
                auto last = matrix[pos.first][pos.second];
                if(cur >= last){
                    visited[x][y] = true;
                    q.push(std::pair<int,int>{x,y});
                }
            }
        }
    }

    vector<vector<int>> pacificAtlantic(vector<vector<int>>& matrix) {
        std::vector<std::vector<int>> res;
        int rows = matrix.size();
        if(rows == 0)
            return res;
        int columns = matrix[0].size();
        vector<vector<bool>> pacific_visited;
        vector<vector<bool>> atlantic_visited;
        pacific_visited.resize(rows);
        atlantic_visited.resize(rows);
        for(int i = 0;i < rows;i++){
            pacific_visited[i].resize(columns);
            atlantic_visited[i].resize(columns);
        }
        std::queue<std::pair<int,int>> pacific_queue;
        std::queue<std::pair<int,int>> atlantic_queue;
        for(int i = 0;i < rows;i++){
            pacific_queue.push(std::pair<int,int>(i,0));
            pacific_visited[i][0] = true;
            atlantic_queue.push(std::pair<int,int>(i,columns - 1));
            atlantic_visited[i][columns - 1] = true;
        }
        for(int i = 1;i < columns;i++){
            pacific_queue.push(std::pair<int,int>(0,i));
            pacific_visited[0][i] = true;
            atlantic_queue.push(std::pair<int,int>(rows - 1,i));
            atlantic_visited[rows - 1][i] = true;
        }
        atlantic_queue.push(std::pair<int,int>(rows - 1,0));
        atlantic_visited[rows - 1][0] = true;
        bfs_pacificAtlantic(matrix,pacific_queue,pacific_visited);
        bfs_pacificAtlantic(matrix,atlantic_queue,atlantic_visited);
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                if(pacific_visited[i][j] && atlantic_visited[i][j]){
                    res.push_back({i,j});
                }
            }
        }
        return res;
    }

    void dfs_pacificAtlantic(vector<vector<int>>& matrix,int row,int col,int pre_height,vector<vector<bool>>& visited){
        auto rows = matrix.size();
        auto columns = matrix[0].size();
        if(row < 0 || row >= rows || col < 0 || col >= columns || visited[row][col] || matrix[row][col] < pre_height)
            return;
        visited[row][col] = true;
        dfs_pacificAtlantic(matrix,row - 1,col,matrix[row][col],visited);
        dfs_pacificAtlantic(matrix,row + 1,col,matrix[row][col],visited);
        dfs_pacificAtlantic(matrix,row,col - 1,matrix[row][col],visited);
        dfs_pacificAtlantic(matrix,row,col + 1,matrix[row][col],visited);
    }

    vector<vector<int>> pacificAtlantic2(vector<vector<int>>& matrix) {
        std::vector<std::vector<int>> res;
        int rows = matrix.size();
        if(rows == 0)
            return res;
        int columns = matrix[0].size();
        vector<vector<bool>> pacific_visited;
        vector<vector<bool>> atlantic_visited;
        pacific_visited.resize(rows);
        atlantic_visited.resize(rows);
        for(int i = 0;i < rows;i++){
            pacific_visited[i].resize(columns);
            atlantic_visited[i].resize(columns);
        }
//        for(int i = 0;i < rows;i++){
//            pacific_visited[i][0] = true;
//            atlantic_visited[i][columns - 1] = true;
//        }
//        for(int i = 0;i < columns;i++){
//            pacific_visited[0][i] = true;
//            atlantic_visited[rows - 1][i] = true;
//        }

        for(int i = 0;i < rows;i++){
            dfs_pacificAtlantic(matrix,i,0,-1,pacific_visited);
            dfs_pacificAtlantic(matrix,i,columns - 1,-1,atlantic_visited);
        }
        for(int i = 0;i < columns;i++){
            dfs_pacificAtlantic(matrix,0,i,-1,pacific_visited);
            dfs_pacificAtlantic(matrix,rows - 1,i,-1,atlantic_visited);
        }
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                if(pacific_visited[i][j] && atlantic_visited[i][j]){
                    res.push_back({i,j});
                }
            }
        }
        return res;
    }
};