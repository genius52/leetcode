#include <vector>
#include <unordered_map>
#include <queue>
#include <map>
using namespace std;

class Solution_743 {
public:
    //djkstra
    int networkDelayTime(vector<vector<int>>& times, int n, int k){
        std::vector<int> distance_from_k(n + 1,2147483647);
        std::unordered_map<int,std::unordered_map<int,int>> graph;
        for(auto t : times){
            graph[t[0]][t[1]] = t[2];
            if(t[0] == k){
                distance_from_k[t[1]] = t[2];
            }
        }
        //first: distance
        //second cur_node
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>, std::greater<std::pair<int,int>>> q;
        std::pair<int,int> start;
        start.first = 0;
        start.second = k;
        q.push(start);
        std::vector<bool> visited(n + 1);
        distance_from_k[k] = 0;
        int max_distance = 0;
        while(!q.empty()){
            auto cur = q.top();
            q.pop();
            if(visited[cur.second])
                continue;
            max_distance = cur.first;
            visited[cur.second] = true;
            if(graph.find(cur.second) == graph.end())
                continue;
            for(auto it = graph[cur.second].begin();it != graph[cur.second].end();it++){
                if((cur.first + (*it).second) > distance_from_k[(*it).first])
                    continue;
                std::pair<int,int> next;
                next.first = cur.first + (*it).second;
                next.second = (*it).first;
                distance_from_k[(*it).first] = next.first;
                q.push(next);
            }
        }

        for (int i = 1; i <= n ; ++i) {
            if(!visited[i])
                return -1;
        }
        return max_distance;
    }

    //bellman ford
//    int networkDelayTime(vector<vector<int>>& times, int n, int k) {
//        std::vector<int> dis(n + 1,2147483647);
//        dis[k] = 0;
//        for(int i = 0;i < n;i++){
//            for(auto t : times){
//                auto cur_start = t[0];
//                auto cur_end = t[1];
//                auto duration = t[2];
//                if(dis[cur_start] != 2147483647 && dis[cur_end] > dis[cur_start] + duration){
//                    dis[cur_end] = dis[cur_start] + duration;
//                }
//            }
//        }
//        int max = 0;
//        for(int i = 1;i <= n;i++){
//            if(dis[i] == 2147483647)
//                return -1;
//            if (dis[i] > max)
//                max = dis[i];
//        }
//        return max;
//    }
};