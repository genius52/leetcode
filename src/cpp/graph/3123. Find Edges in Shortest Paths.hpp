#include <vector>
#include <queue>
using namespace std;

class Solution_3123 {
public:
    vector<bool> findAnswer(int n, vector<vector<int>>& edges) {
        std::vector<std::vector<std::pair<int,int>>> graph(n);
        for(auto edge : edges){
            graph[edge[0]].push_back({edge[1],edge[2]});
        }
        std::vector<int> dis_from_zero(n);
        for(int i = 1;i < n;i++){
            dis_from_zero[i] = 1 << 31 - 1;
        }
        //distance,node
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        std::vector<bool> res(n);
        q.push({0,0});
        while (!q.empty()){
            auto top = q.top();
            q.pop();
        }
        return res;
    }
};