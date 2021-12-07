#include <vector>
#include <queue>
using namespace std;

class Solution_1786 {
public:
    int dfs_countRestrictedPaths(int cur,int n,std::vector<std::vector<std::pair<int,int>>>& graph,std::vector<int>& disN,
                                 std::vector<int>& memo){
        if(cur == n)
            return 1;
        if(memo[cur] != -1){
            return memo[cur];
        }
        int res = 0;
        for(int i = 0;i < graph[cur].size();i++){
            int neighbour = graph[cur][i].first;
            if(disN[cur] > disN[neighbour]){
                res += dfs_countRestrictedPaths(neighbour,n,graph,disN,memo);
                res %= int(1e9 + 7);
            }
        }
        return memo[cur] = res;
    }

    int countRestrictedPaths(int n, vector<vector<int>>& edges) {
        std::vector<int> disN(n + 1);
        for(int i = 1;i <= n - 1;i++)
            disN[i] = 2147483647;
        std::vector<std::vector<std::pair<int,int>>> graph(n + 1);
        for(auto edge : edges){
            graph[edge[0]].push_back({edge[1],edge[2]});
            graph[edge[1]].push_back({edge[0],edge[2]});
        }
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        q.push({0,n});
        std::vector<bool> visited(n + 1);
        int MOD = 1e9 + 7;
        while(!q.empty()){
            auto cur = q.top();
            q.pop();
            if(visited[cur.second])
                continue;
            if(cur.first > disN[cur.second])
                continue;
            visited[cur.second] = true;
            for(int i = 0;i < graph[cur.second].size();i++){
                int neighbour = graph[cur.second][i].first;
                int weight = graph[cur.second][i].second;
                if((cur.first + weight) < disN[neighbour]){
                    disN[neighbour] = cur.first + weight;
                    q.push({disN[neighbour],neighbour});
                }
            }
        }
        std::vector<int> memo(n + 1);
        for (int i = 1; i <= n; ++i) {
            memo[i] = -1;
        }
        dfs_countRestrictedPaths(1,n,graph,disN,memo);
        return memo[1];
    }
};