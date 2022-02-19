#include <vector>
#include <queue>
using namespace std;

class Solution_2050 {
public:
    int minimumTime(int n, vector<vector<int>>& relations, vector<int>& time) {
        std::vector<std::vector<int>>  graph(n + 1);
        std::vector<int>  indegree(n + 1);
        std::vector<int>  finish_time(n + 1,0);
        int res = 0;
        for(auto edge : relations){
            indegree[edge[1]]++;
            graph[edge[0]].push_back(edge[1]);
        }
        //std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;//small top
        std::queue<std::pair<int,int>> q;
        for(int i = 1;i <= n;i++){
            if(indegree[i] == 0){
                q.push({time[i - 1],i});
                finish_time[i] = time[i - 1];
                //start_time[i] = 0;
            }
            res = max(res, time[i - 1]);
        }
        while (!q.empty()){
            int cur_len = q.size();
            for(int i = 0;i < cur_len;i++){
                auto cur = q.front();
                q.pop();
                int duration = cur.first;
                for(auto next : graph[cur.second]){
                    indegree[next]--;
                    finish_time[next] = max(finish_time[next],duration + time[next - 1]);
                    if(indegree[next] == 0){
                        q.push({finish_time[next],next});
                        res = max(res,finish_time[next]);
                    }
                }
            }
        }
        return res;
    }
};