#include <vector>
#include <map>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution_1584 {
public:
    int minCostConnectPoints(vector<vector<int>>& points) {
        int l = points.size();
        if (l <= 1){
            return 0;
        }
        std::unordered_map<int,std::map<int,int>> graph;
        for (int i = 0;i < l;i++){
            for (int j = 0;j < l;j++){
                if (i == j)
                    continue;
                graph[i][j] = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
            }
        }
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        int res = 0;
        std::vector<bool> visited(l);
        int cnt = 1;
        int node = 0;
        while(cnt < l){
            visited[node] = true;
            for (int i = 0;i < l;i++){
                if (i == node || visited[i])
                    continue;
                std::pair<int,int> p(graph[node][i],i);
                q.push(p);
            }
            while(visited[q.top().second])
                q.pop();
            res += q.top().first;
            node = q.top().second;
            q.pop();
            cnt++;
        }
        return res;
    }
};