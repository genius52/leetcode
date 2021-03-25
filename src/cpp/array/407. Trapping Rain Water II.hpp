#include <vector>
#include <queue>
using namespace std;

//Given the following 3x6 height map:
//[
//  [1,4,3,1,3,2],
//  [3,2,1,3,2,4],
//  [2,3,3,2,3,1]
//]
//
//Return 4.
class Solution_407 {
public:
    int trapRainWater(vector<vector<int>>& heightMap) {
        int rows = heightMap.size();
        int columns = heightMap[0].size();
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;
        for(int i = 0;i < rows;i++){
            std::pair<int,std::pair<int,int>> p1,p2;
            p1.first = heightMap[i][0];
            p1.second.first = i;
            p1.second.second = 0;
            p2.first = heightMap[i][columns - 1];
            p2.second.first = i;
            p2.second.second = columns - 1;
            q.push(p1);
            q.push(p2);
            visited[i][0] = true;
            visited[i][columns - 1] = true;
        }
        for(int j = 1;j < columns - 1;j++){
            std::pair<int,std::pair<int,int>> p1,p2;
            p1.first = heightMap[0][j];
            p1.second.first = 0;
            p1.second.second = j;
            p2.first = heightMap[rows - 1][j];
            p2.second.first = rows - 1;
            p2.second.second = j;
            q.push(p1);
            q.push(p2);
            visited[0][j] = true;
            visited[rows - 1][j] = true;
        }
        int total = 0;
        int dir[4][4] = {{-1,0},{1,0},{0,-1},{0,1}};
        int lowest = 0;
        while (!q.empty()){
            int len = q.size();

            for(int i = 0;i < len;i++){
                auto node = q.top();
                q.pop();
                lowest = max(lowest,node.first);
                for(int j = 0;j < 4;j++){
                    int r = node.second.first + dir[j][0];
                    int c = node.second.second + dir[j][1];
                    if(r >= 0 && r < rows && c >= 0 && c < columns && !visited[r][c] ){
                        visited[r][c] = true;
                        if(heightMap[r][c] < lowest){
                            total += lowest - heightMap[r][c];//???
                        }
                        std::pair<int,std::pair<int,int>> p;
                        p.first = heightMap[r][c];
                        p.second.first = r;
                        p.second.second = c;
                        q.push(p);
                    }
                }
            }
        }
        return total;
    }
};