#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

//Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//Output: 65
//Explanation: robot will be stuck at (1, 4) before turning left and going to (1, 8)
class Solution_874 {
public:
    int robotSim(vector<int>& commands, vector<vector<int>>& obstacles) {
        std::unordered_map<int,std::unordered_set<int>> graph;
        for(auto it = obstacles.begin();it != obstacles.end();it++){
            graph[(*it)[0]].insert((*it)[1]);
        }
        int max_square = 0;
        std::vector<std::vector<int>> dirs{{0,1},{1,0},{0,-1},{-1,0}};
        int cmd_len = commands.size();
        //std::vector<std::vector<int>> position;
        //position.resize(cmd_len);
        int dir = 0;
        int cur_x = 0;
        int cur_y = 0;
        for(int i = 0;i < cmd_len;i++){
            if(commands[i] == -2){
                dir = (dir + 3) % 4;
            }else if(commands[i] == -1){
                dir = (dir + 1) % 4;
            }else{
                for(int j = 0;j < commands[i];j++){
                    cur_x += dirs[dir][0];
                    cur_y += dirs[dir][1];
                    if(graph.find(cur_x) != graph.end()){
                        if(graph[cur_x].find(cur_y) != graph[cur_x].end()){
                            cur_x -= dirs[dir][0];
                            cur_y -= dirs[dir][1];
                            break;
                        }
                    }
                }
                int s = cur_x * cur_x + cur_y * cur_y;
                if(s > max_square)
                    max_square = s;
            }
        }
        return max_square;
    }
};