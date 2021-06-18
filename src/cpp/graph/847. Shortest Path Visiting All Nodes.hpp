#include <vector>
#include <unordered_set>
#include <string>
#include <queue>
using namespace std;

class Solution_847 {
public:
    int shortestPathLength(vector<vector<int>>& graph) {
        int n = graph.size();
        if(n <= 1)
            return 0;
        //（curNode | VisitedNode）表示状态
        std::queue<std::pair<int,int>> q;//curnode << 16 | state   first:当前节点，second:状态
        std::unordered_set<std::string> visited;
        int target = 0;
        for(int i = 0;i < n;i++){
            int mask = 1 << i;
            target |= mask;
            q.push({i,mask});
            std::string key = std::to_string(i) + "-" + std::to_string(mask);
            visited.insert(key);
        }
        int steps = 0;
        while(q.size() > 0){
            int l = q.size();
            for(int i = 0;i < l;i++){
                std::pair<int,int> cur = q.front();
                q.pop();
                if(cur.second == target)
                    return steps;
                //visited.insert(std::to_string(cur.first) + "-" + std::to_string(cur.second));
                std::vector<int> neighbours = graph[cur.first];
                for(int j = 0;j < neighbours.size();j++){
                    std::pair<int,int> next;
                    next.first = graph[cur.first][j];
                    next.second = (1 << next.first) | cur.second;
                    std::string key = std::to_string(next.first) + "-" + std::to_string(next.second);
                    if(visited.find(key) != visited.end())
                        continue;
                    visited.insert(key);
                    q.push(next);
                }
            }
            steps++;
        }
        return -1;
    }
};