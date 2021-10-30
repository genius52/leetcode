#include <vector>
#include <queue>
using namespace std;

class Solution_1514 {
public:
    //bfs, djistra
    double maxProbability(int n, vector<vector<int>>& edges, vector<double>& succProb, int start, int end) {
        std::vector<std::unordered_map<int,double>> graph(n);
        int len = edges.size();
        for (int i = 0; i < len; ++i) {
            graph[edges[i][0]][edges[i][1]] = succProb[i];
            graph[edges[i][1]][edges[i][0]] = succProb[i];
        }
        std::vector<bool> visited(n);
        std::vector<double> start_to_cur(n);
        start_to_cur[start] = 1.0;
        std::priority_queue<std::pair<double,int>> q;
        q.push({1.0,start});
        //double res = 0;
        while (!q.empty()){
            auto cur = q.top();
            q.pop();
            if(cur.second == end){
                //res = max(res,cur.first);
                continue;
            }
            for (auto it : graph[cur.second]) {
                if(it.second != 0){
                    double rate = cur.first * it.second;
                    if (start_to_cur[it.first] < rate){
                        start_to_cur[it.first] = rate;
                        q.push({rate,it.first});
                    }
                }
            }
        }
        return start_to_cur[end];
    }
};