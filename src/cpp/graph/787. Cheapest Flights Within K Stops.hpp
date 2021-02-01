#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
class Solution_787 {
public:
//    int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst, int K) {
//        std::vector<int> from_src(n,1e8);
//        from_src[src] = 0;
//        for(int i = 0;i <= K;i++){
//            std::vector<int> copy = from_src;
//            for(auto f : flights){
//                auto cur_src = f[0];
//                auto cur_dst = f[1];
//                auto cur_cost = f[2];
//                if(copy[cur_dst] > from_src[cur_src] + cur_cost){
//                    copy[cur_dst] = from_src[cur_src] + cur_cost;
//                }
//            }
//            from_src = copy;
//        }
//        if(from_src[dst] == 1e8){
//            return -1;
//        }
//        return from_src[dst];
//    }

    int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst, int K) {
        std::unordered_map<int,std::unordered_map<int,int>> graph;
        int len = flights.size();
        for(int i = 0;i < len;i++){
            graph[flights[i][0]][flights[i][1]] = flights[i][2];
        }
        std::queue<std::pair<int,int>> q;
        q.push({src,0});//min cost from source to cur node
        int res = 2147483647;
        while(!q.empty() && K + 1 >= 0){
            int l = q.size();
            for(int i = 0;i < l;i++){
                auto cur = q.front();
                q.pop();
                if(cur.first == dst){
                    res = min(res,cur.second);
                }
                if(graph.find(cur.first) == graph.end())
                    continue;
                for(auto it = graph[cur.first].begin();it != graph[cur.first].end();it++){
                    if(cur.second + it->second > res)
                        continue;
                    q.push({it->first,cur.second + it->second});
                }
            }
            K--;
        }
        if(res == 2147483647)
            return -1;
        return res;
    }
};