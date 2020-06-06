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
//    vector<vector<int>> pacificAtlantic(vector<vector<int>>& matrix) {
//        std::vector<std::vector<int>> res;
//        int rows = matrix.size();
//        if(rows == 0)
//            return res;
//        int columns = matrix[0].size();
//        std::vector<std::vector<bool>> dp_left_top;
//        std::vector<std::vector<bool>> dp_right_bottom;
//        dp_left_top.resize(rows);
//        dp_right_bottom.resize(rows);
//        for(int i = 0;i < rows;i++){
//            dp_left_top[i].resize(columns);
//            dp_right_bottom[i].resize(columns);
//        }
//        for(int i = 0;i < rows;i++){
//            for(int j = 0;j < columns;j++){
//                if(i == 0 || j == 0){
//                    dp_left_top[i][j] = true;
//                    continue;
//                }
//                if((matrix[i][j] >= matrix[i - 1][j] && dp_left_top[i - 1][j]) || (matrix[i][j] >= matrix[i][j - 1] && dp_left_top[i][j - 1]))
//                    dp_left_top[i][j] = true;
//            }
//        }
//        for(int i = rows - 1;i >= 0;i--){
//            for(int j = columns - 1;j >= 0;j--){
//                if(i == rows - 1 || j == columns - 1){
//                    dp_right_bottom[i][j] = true;
//                    continue;
//                }
//                if((matrix[i][j] >= matrix[i + 1][j] && dp_right_bottom[i+1][j]) || (matrix[i][j] >= matrix[i][j + 1] && dp_right_bottom[i][j+1]))
//                    dp_right_bottom[i][j] = true;
//            }
//        }
//        for(int i = 0;i < rows;i++) {
//            for (int j = 0; j < columns; j++) {
//                if(dp_left_top[i][j] && dp_right_bottom[i][j])
//                    res.push_back(std::vector<int>{i,j});
//            }
//        }
//        return res;
//    }
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
};