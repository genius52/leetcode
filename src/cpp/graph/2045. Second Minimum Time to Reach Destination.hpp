#include <vector>
#include <queue>
using namespace std;

class Solution_2045 {
public:
    int secondMinimum(int n, vector<vector<int>>& edges, int time, int change) {
        std::vector<std::vector<int>> graph(n + 1);
        for(auto edge : edges){
            graph[edge[0]].push_back(edge[1]);
            graph[edge[1]].push_back(edge[0]);
        }
        std::vector<int> dis1(n + 1,2147483647);
        dis1[1] = 0;
        std::vector<int> dis2(n + 1,2147483647);
        //dis2[1] = 0;
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        q.push({0,1});
        while(!q.empty()){
            auto p = q.top();
            q.pop();
//            int next_duration = p.first;
//            bool is_red = p.first % change % 2;
//            if(is_red){
//                next_duration = p.first / change * change + change + time;
//            }else{
//                next_duration += time;
//            }
            int next_duration = 0;
            int times = p.first /change;
            if(times % 2 == 1){
                next_duration =  (times+1)*change + time;
            }else{
                next_duration = p.first + time;
            }
            int node = p.second;
            for (auto neighbour : graph[node]) {
                if(next_duration < dis1[neighbour]){
                    dis2[neighbour] = dis1[neighbour];
                    dis1[neighbour] = next_duration;
                    q.push({next_duration,neighbour});
                }else if(next_duration > dis1[neighbour] && next_duration < dis2[neighbour]){
                    dis2[neighbour] = next_duration;
                    q.push({next_duration,neighbour});
                }
            }
        }
        return dis2[n];
    }
};