#include <vector>
#include <map>
#include <unordered_map>
#include <unordered_set>
#include <queue>
using namespace std;

class Solution_310 {
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges){
        if(n == 1){
            return std::vector<int>{0};
        }
        std::unordered_map<int,std::unordered_set<int>> graph;
        for(int i = 0;i < edges.size();i++){
            graph[edges[i][0]].insert(edges[i][1]);
            graph[edges[i][1]].insert(edges[i][0]);
        }
        std::queue<int> q;
        //std::vector<bool> visited(n);
        for(auto it = graph.begin();it != graph.end();it++){
            if((*it).second.size() == 1){
                q.push((*it).first);
                //visited[(*it).first] = true;
            }
        }
        while(n > 2){
            int len = q.size();
            n -= len;
            for(int i = 0;i < len;i++){
                auto n = q.front();
                q.pop();
                for(auto it = graph[n].begin();it != graph[n].end();it++){
                    int neighbour = *it;
                    graph[neighbour].erase(n);
                    if(graph[neighbour].size() == 1){
                        q.push(*it);
                    }
                }
            }
        }
        std::vector<int> res;
        while(!q.empty()){
            res.push_back(q.front());
            q.pop();
        }
        return res;
    }
//    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges){
//        std::unordered_map<int,std::vector<int>> graph;
//        for(int i = 0;i < edges.size();i++){
//            graph[edges[i][0]].push_back(edges[i][1]);
//            graph[edges[i][1]].push_back(edges[i][0]);
//        }
//        std::map<int,std::vector<int>> record;
//        int min_step = 2147483647;
//        for(int i = 0;i < n;i++){
//            std::queue<int> q;
//            q.push(i);
//            int steps = 0;
//            std::vector<bool> visited(n);
//            while(q.size() > 0){
//                int len = q.size();
//                for(int j = 0;j < len;j++){
//                    auto n = q.front();
//                    visited[n] = true;
//                    q.pop();
//                    for(auto it = graph[n].begin();it != graph[n].end();it++){
//                        if(visited[*it]){
//                            continue;
//                        }
//                        q.push(*it);
//                    }
//                }
//                steps++;
//                if(steps > min_step){
//                    break;
//                }
//            }
//            if (steps < min_step){
//                min_step = steps;
//            }
//            record[steps].push_back(i);
//        }
//        return record.begin()->second;
//    }

//    int dfs_findMinHeightTrees(std::unordered_map<int,std::vector<int>>& graph,std::vector<bool>& visited,int node_num,int height,int visited_num,int target_num,int pre_min){
//        if(height > pre_min)
//            return height;
//        if(visited_num == target_num)
//            return height;
//        if(visited[node_num])
//            return height;
//        visited[node_num] = true;
//        visited_num++;
//        int max_height = height;
//        for(int i = 0;i < graph[node_num].size();i++){
//            if(visited[graph[node_num][i]])
//                continue;
//            int h = dfs_findMinHeightTrees(graph,visited,graph[node_num][i],height + 1,visited_num,target_num,pre_min);
//            if(h > max_height)
//                max_height = h;
//        }
//        return max_height;
//    }
//
//    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
//        std::vector<int> res;
//        std::unordered_map<int,std::vector<int>> graph;
//        for(int i = 0;i < edges.size();i++){
//            graph[edges[i][0]].push_back(edges[i][1]);
//            graph[edges[i][1]].push_back(edges[i][0]);
//        }
//        int min_height = n;
//        std::unordered_map<int,std::vector<int>> record;
//        std::vector<bool> visited(n,false);
//        for(int i = 0;i < n;i++){
//            int height = dfs_findMinHeightTrees(graph,visited,i,0,0,n,min_height);
//            if(height < min_height)
//                min_height = height;
//            record[height].push_back(i);
//            visited.assign(n,false);
//        }
//        return record[min_height];
//    }
};