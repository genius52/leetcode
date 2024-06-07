#include <vector>
#include <queue>
using namespace std;

class Solution_3123 {
public:
    void dfs_findAnswer(int cur,std::vector<std::vector<std::vector<int>>>& graph,std::vector<int>& dis_from_zero, std::vector<bool>& visited,std::vector<bool>& res){
        visited[cur] = true;
        for(auto edge : graph[cur]){
            auto next = edge[0];
            auto weight = edge[1];
            auto idx = edge[2];
            if(dis_from_zero[next] + weight != dis_from_zero[cur])
                continue;
            if(!visited[next]){
                res[idx] = true;
                dfs_findAnswer(next,graph,dis_from_zero,visited,res);
            }
        }
    }

    vector<bool> findAnswer(int n, vector<vector<int>>& edges) {
        std::vector<std::vector<std::vector<int>>> graph(n);
        for(int i = 0;i < edges.size();i++){
            graph[edges[i][0]].push_back({edges[i][1],edges[i][2],i});
            graph[edges[i][1]].push_back({edges[i][0],edges[i][2],i});
        }
        std::vector<int> dis_from_zero(n);
        for(int i = 1;i < n;i++){
            dis_from_zero[i] = 1 << 31 - 1;
        }
        //distance,node
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        std::vector<bool> res(edges.size());
        q.push({0,0});//距离，节点
        while (!q.empty()){
            auto top = q.top();
            q.pop();
            if(top.first > dis_from_zero[top.second])
                continue;
            for(auto next:graph[top.second]){
                if(top.first + next[1] < dis_from_zero[next[0]]){
                    q.push({top.first + next[1],next[0]});
                    dis_from_zero[next[0]] = top.first + next[1];
                }
            }
        }
        if(dis_from_zero[n - 1] == 1 << 31 - 1)
            return res;
        std::vector<bool> visited(n);
        dfs_findAnswer(n - 1,graph,dis_from_zero,visited,res);
        return res;
    }
};