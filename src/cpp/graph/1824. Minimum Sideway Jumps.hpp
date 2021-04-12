#include <vector>
using namespace std;

class Solution_1824 {
public:
    int dfs_minSideJumps(vector<int>& obstacles,int len,int cur_pos,int cur_road,std::vector<std::vector<int>>& visited){
        if(cur_pos >= len){
            return 0;
        }
        if(visited[cur_pos][cur_road - 1] != 0){
            return visited[cur_pos][cur_road - 1];
        }

        int next = cur_pos + 1;
        int min_steps = 2147483647;
        if(obstacles[next] == cur_road){//block on next road
            for(int i = 1;i <= 3;i++){
                if(i == cur_road){
                    continue;
                }
                if(obstacles[cur_pos] == i)
                    continue;
                int ret = dfs_minSideJumps(obstacles,len,cur_pos,i,visited);
                if(ret == 2147483647)
                    continue;
                int steps = 1 + ret;
                if(steps < min_steps)
                    min_steps = steps;
            }
        }else{
            min_steps = dfs_minSideJumps(obstacles,len,cur_pos + 1,cur_road,visited);
        }
        visited[cur_pos][cur_road - 1] = min_steps;
        return min_steps;
    }
    int minSideJumps(vector<int>& obstacles) {
        int len = obstacles.size();
        std::vector<std::vector<int>> visited(len,std::vector<int>(3));
        return dfs_minSideJumps(obstacles,len,0,2,visited);
    }
};