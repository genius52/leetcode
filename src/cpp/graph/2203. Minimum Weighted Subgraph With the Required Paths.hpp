#include <vector>
#include <queue>
using namespace std;

class Solution_2203 {
public:
    long long minimumWeight(int n, vector<vector<int>>& edges, int src1, int src2, int dest) {
        std::vector<std::vector<std::pair<int,int>>> graph(n);
        for(auto edge : edges){
            graph[edge[0]].push_back({edge[1],edge[2]});
        }
        std::vector<long long> from_src1(n,LLONG_MAX);
        from_src1[src1] = 0;
        //std::vector<bool> visited1(n);
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>,std::greater<std::pair<long long,int>>> q1;
        q1.push({0,src1});
        while (!q1.empty()){
            auto cur = q1.top();
            long long duration = cur.first;
            int node = cur.second;
            q1.pop();
            if(duration > from_src1[node])
                continue;
            for(int i = 0;i < graph[node].size();i++){
                int next_node = graph[node][i].first;
                int weight = graph[node][i].second;
                if(duration + weight < from_src1[next_node]){
                    from_src1[next_node] = duration + weight;
                    q1.push({duration + weight,next_node});
                }
            }
        }
        if(from_src1[dest] == LLONG_MAX)
            return -1;
        std::vector<long long> from_src2(n,LLONG_MAX);
        from_src2[src2] = 0;
        //std::vector<bool> visited2(n);
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>,std::greater<std::pair<long long,int>>> q2;
        q2.push({0,src2});
        while (!q2.empty()){
            auto cur = q2.top();
            long long duration = cur.first;
            int node = cur.second;
            q2.pop();
            if(duration > from_src2[node])
                continue;
            for(int i = 0;i < graph[node].size();i++){
                int next_node = graph[node][i].first;
                int weight = graph[node][i].second;
                if(duration + weight < from_src2[next_node]){
                    from_src2[next_node] = duration + weight;
                    q2.push({duration + weight,next_node});
                }
            }
        }
        if(from_src2[dest] == LLONG_MAX)
            return -1;
        std::vector<std::vector<std::pair<int,int>>> graph2(n);
        for(auto edge : edges){
            graph2[edge[1]].push_back({edge[0],edge[2]});
        }
        std::vector<long long> from_dest(n,LLONG_MAX);
        from_dest[dest] = 0;
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>,std::greater<std::pair<long long,int>>> q3;
        q3.push({0,dest});
        while (!q3.empty()){
            auto cur = q3.top();
            long long duration = cur.first;
            int node = cur.second;
            q3.pop();
            if(duration > from_dest[node])
                continue;
            for(int i = 0;i < graph2[node].size();i++){
                int next_node = graph2[node][i].first;
                int weight = graph2[node][i].second;
                if(duration + weight < from_dest[next_node]){
                    from_dest[next_node] = duration + weight;
                    q3.push({duration + weight,next_node});
                }
            }
        }
        long long res = LLONG_MAX;
        for(int i = 0;i < n;i++){
            if (from_src1[i] != LLONG_MAX && from_src2[i] != LLONG_MAX && from_dest[i] != LLONG_MAX)
                res = min(res,from_src1[i] + from_src2[i] + from_dest[i]);
        }
        if(res == LLONG_MAX)
            return -1;
        return res;
    }
};