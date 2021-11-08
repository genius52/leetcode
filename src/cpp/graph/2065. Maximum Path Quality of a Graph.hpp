#include <vector>
#include <unordered_set>
#include <set>
#include <unordered_map>
#include <queue>
using namespace std;

//输入：values = [0,32,10,43], edges = [[0,1,10],[1,2,15],[0,3,10]], maxTime = 49
//输出：75
//解释：
//一条可能的路径为：0 -> 1 -> 0 -> 3 -> 0 。总花费时间为 10 + 10 + 10 + 10 = 40 <= 49 。
//访问过的节点为 0 ，1 和 3 ，最大路径价值为 0 + 32 + 43 = 75
//class Solution_2065 {
//public:
//    void dfs(std::vector<std::unordered_map<int,int>>& graph,std::unordered_set<int>& visited,int cur,int& max_res){
//
//    }
//
//    int maximalPathQuality(vector<int>& values, vector<vector<int>>& edges, int maxTime) {
//        int len = values.size();
//        std::vector<std::unordered_map<int,int>> graph(len);
//        for (int i = 0; i < edges.size(); ++i) {
//            graph[edges[i][0]][edges[i][1]] = edges[i][2];
//            graph[edges[i][1]][edges[i][0]] = edges[i][2];
//        }
//        int res = 0;
//        std::unordered_set<int> visited;
//        dfs(graph,visited,0,res);
//        return res;
//    }
//};

//TLE
class Solution_2065 {
public:
    struct Trace{
        int cost = 0;
        int cur_node = 0;
        int value = 0;
        std::unordered_set<int> visited;
    };
    int maximalPathQuality(vector<int>& values, vector<vector<int>>& edges, int maxTime) {
        int len = values.size();
        std::vector<std::unordered_map<int,int>> graph(len);
        for (int i = 0; i < edges.size(); ++i) {
            graph[edges[i][0]][edges[i][1]] = edges[i][2];
            graph[edges[i][1]][edges[i][0]] = edges[i][2];
        }
        std::queue<Trace> q;
        Trace first;
        first.cur_node = 0;
        first.cost = 0;
        first.value = values[0];
        first.visited.insert(0);
        q.push(first);
        int max_val = values[0];
        while (q.size() > 0){
            Trace cur = q.front();
            q.pop();
            int cur_node = cur.cur_node;
            for(auto it = graph[cur_node].begin();it != graph[cur_node].end();it++){
                int nextcost = cur.cost + (*it).second;
                if(nextcost > maxTime)
                    continue;
                Trace next = cur;
                next.cost = nextcost;
                next.cur_node = (*it).first;
                next.visited = cur.visited;
                next.value = cur.value;
                if(next.visited.find(next.cur_node) == next.visited.end()){
                    next.value += values[next.cur_node];
                    next.visited.insert(next.cur_node);
                }
                if(next.cur_node == 0){
                    max_val = max(max_val,next.value);
                }
                q.push(next);
            }
        }
        return max_val;
    }
};