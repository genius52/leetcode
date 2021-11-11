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
        std::vector<std::vector<int>> graph(l,std::vector<int>(l));
        for (int i = 0;i < l;i++){
            for (int j = i + 1;j < l;j++){
                if (i == j)
                    continue;
                graph[i][j] = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
                graph[j][i] = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
            }
        }
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        for(int i = 0;i < l;i++){
            if(graph[0][i] != 0)
                q.push({graph[0][i],i});//把0的所有边加入队列
        }
        std::vector<bool> visited(l);
        visited[0] = true;
        int cnt = 1;
        int res = 0;
        while(!q.empty() && cnt < l){
            while (!q.empty() && visited[q.top().second]){
                q.pop();
            }
            auto cur = q.top();
            res += cur.first;
            visited[cur.second] = true;
            cnt++;
            for(int i = 0;i < l;i++){
                if(graph[cur.second][i] != 0 && visited[cur.second])
                    q.push({graph[cur.second][i],i});
            }
        }
        return res;
    }

//    int minCostConnectPoints(vector<vector<int>>& points) {
//        int l = points.size();
//        if (l <= 1){
//            return 0;
//        }
//        std::unordered_map<int,std::map<int,int>> graph;
//        for (int i = 0;i < l;i++){
//            for (int j = 0;j < l;j++){
//                if (i == j)
//                    continue;
//                graph[i][j] = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
//            }
//        }
//        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
//        int res = 0;
//        std::vector<bool> visited(l);
//        int cnt = 1;
//        int node = 0;
//        while(cnt < l){
//            visited[node] = true;
//            for (int i = 0;i < l;i++){
//                if (i == node || visited[i])
//                    continue;
//                std::pair<int,int> p(graph[node][i],i);
//                q.push(p);
//            }
//            while(visited[q.top().second])
//                q.pop();
//            res += q.top().first;
//            node = q.top().second;
//            q.pop();
//            cnt++;
//        }
//        return res;
//    }
};