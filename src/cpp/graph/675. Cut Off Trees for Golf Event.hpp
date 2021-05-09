#include <vector>
#include <queue>
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution_675 {
public:
    int src_dst(vector<vector<int>>& forest,int start_x,int start_y,int dst_x,int dst_y){
        int rows = forest.size();
        int columns = forest[0].size();
        if(start_x == dst_x && start_y == dst_y)
            return 0;
        int dir[4][2] = {{-1,0},{0,1},{1,0},{0,-1}};
        std::queue<std::pair<int,int>> q;
        q.push(std::pair<int,int>{start_x,start_y});
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        visited[start_x][start_y] = true;
        int steps = 1;
        while(!q.empty()){
            int len = q.size();
            for(int i = 0;i < len;i++){
                auto f = q.front();
                q.pop();
                for(int j = 0;j < 4;j++){
                    int x = f.first + dir[j][0];
                    int y = f.second + dir[j][1];
                    if(x == dst_x && y == dst_y)
                        return steps;
                    if(x >= 0 && x < rows && y >= 0 && y < columns && forest[x][y] != 0 && !visited[x][y]){
                        visited[x][y] = true;
                        q.push(std::pair<int,int>{x,y});
                    }
                }
            }
            steps++;
        }
        return -1;
    }

    int cutOffTree(vector<vector<int>>& forest) {
        int rows = forest.size();
        int columns = forest[0].size();
        std::vector<std::vector<int>> height_pos;
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                if(forest[i][j] > 1)
                    height_pos.push_back({forest[i][j],i,j});
            }
        }
        std::sort(height_pos.begin(),height_pos.end());
        int start_x = 0;
        int start_y = 0;
        int target_index = 0;
        int len = height_pos.size();
        int steps = 0;
        while(target_index < len){
            int dst_x = height_pos[target_index][1];
            int dst_y = height_pos[target_index][2];
            target_index++;
            if(start_x == dst_x && start_y == dst_y)
                continue;
            auto distance = src_dst(forest,start_x,start_y,dst_x,dst_y);
            if(distance == -1)
                return -1;
            steps += distance;
            start_x = dst_x;
            start_y = dst_y;
        }
        return steps;
    }
};