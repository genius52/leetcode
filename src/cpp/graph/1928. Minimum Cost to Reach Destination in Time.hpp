#include <vector>
#include <queue>
using namespace std;

class Solution_1928 {
public:
    int minCost(int maxTime, vector<vector<int>>& edges, vector<int>& passingFees) {
        int len = passingFees.size();
        std::vector<std::vector<std::pair<int,int>>> graph(len);
        for(auto edge : edges){
            graph[edge[0]].push_back(std::make_pair(edge[1],edge[2]));
            graph[edge[1]].push_back(std::make_pair(edge[0],edge[2]));
        }
        std::priority_queue<std::vector<int>,vector<std::vector<int>>,greater<std::vector<int>>>  q;//fee,time,node
        std::vector<int> duration_from_start(len,2147483647);
        duration_from_start[0] = 0;
        std::vector<int> fee_from_start(len,2147483647);
        fee_from_start[0] = 0;
        q.push({passingFees[0],0,0});
        while (!q.empty()){
            auto cur = q.top();
            q.pop();
            int cur_fee = cur[0];
            int cur_duration = cur[1];
            int cur_node = cur[2];
            if(cur_node == len - 1)
                return cur_fee;
            for(int i = 0;i < graph[cur_node].size();i++){
                int neighbour_node = graph[cur_node][i].first;
                int next_duration = graph[cur_node][i].second;
                if(cur_duration + next_duration > maxTime){
                    continue;
                }
                if((cur_duration + next_duration < duration_from_start[neighbour_node]) || ((cur_fee + passingFees[neighbour_node]) < fee_from_start[neighbour_node])){
                    duration_from_start[neighbour_node] = cur_duration + next_duration;
                    fee_from_start[neighbour_node] = cur_fee + passingFees[neighbour_node];
                    q.push({cur_fee + passingFees[neighbour_node],cur_duration + next_duration,neighbour_node});
                }
            }
        }
        if(fee_from_start[len - 1] == 2147483647)
            return -1;
        return fee_from_start[len - 1];
    }
};