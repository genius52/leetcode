#include <vector>
#include <queue>
#include <unordered_map>
using namespace std;

class Solution_882 {
public:
    int reachableNodes(vector<vector<int>>& edges, int maxMoves, int n) {
        std::unordered_map<int,std::unordered_map<int,int>> graph;
        for(auto edge : edges){
            graph[edge[0]][edge[1]] = edge[2];
            graph[edge[1]][edge[0]] = edge[2];
        }
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>, std::less<std::pair<int,int>>> q;
        std::vector<bool> visited(n);
        q.push({maxMoves,0});
        int res = 0;
        while(q.size() > 0){
            std::pair<int,int> cur = q.top();//每次选步数最多的节点
            q.pop();
            if(visited[cur.second]){
                continue;
            }
            visited[cur.second] = true;
            res++;
            for(auto neighbour : graph[cur.second]){
                int cost = 1 + neighbour.second;
                if(cost > 0){
                    if(cur.first >= cost && !visited[neighbour.first]){//有多余步数加入队列
                        q.push({cur.first - cost,neighbour.first});
                    }
                    int reach = min(cur.first, cost - 1);
                    graph[neighbour.first][cur.second] -= reach;
                    //graph[neighbour.first][cur.second] -= min(cur.first,neighbour.second);
                    res += reach;
                }
            }
        }
        return res;
    }
};