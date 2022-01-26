#include <vector>
#include <queue>
using namespace std;

class Solution_1976 {
public:
    int countPaths(int n, vector<vector<int>>& roads) {
        std::vector<std::vector<std::pair<int,int>>> graph(n);
        for(auto road : roads){
            graph[road[0]].push_back(std::make_pair(road[1],road[2]));
            graph[road[1]].push_back(std::make_pair(road[0],road[2]));
        }
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>,std::greater<std::pair<long long,int>>> q;
        q.push({0,0});
        std::vector<long long> from_start(n,9223372036854775807);
        from_start[0] = 0;
        std::vector<long long> cnt(n,0);
        cnt[0] = 1;
        int res = 0;
        int MOD = 1e9 + 7;
        while(!q.empty()){
            auto top = q.top();
            long long cur_cost = top.first;
            int src = top.second;
            q.pop();
            if(src == n - 1)
                continue;
            for(int i = 0;i < graph[src].size();i++){
                int next_node = graph[src][i].first;
                long next_cost = graph[src][i].second;
                if((long long)(cur_cost) + (long long)(next_cost) < from_start[next_node]){
                    from_start[next_node] = (long long)(cur_cost) + (long long)(next_cost);
                    q.push(std::make_pair(from_start[next_node],next_node));
                    cnt[next_node] = cnt[src];
                    cnt[next_node] %= MOD;
                }else if((long long)(cur_cost) + (long long)(next_cost)  == from_start[next_node]){
                    cnt[next_node] += cnt[src];
                    cnt[next_node] %= MOD;
                }
            }
        }
        return cnt[n - 1];
    }
};