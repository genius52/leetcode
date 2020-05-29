#include <vector>
#include <unordered_map>
using namespace std;

class Solution_310 {
public:
    int dfs_findMinHeightTrees(std::unordered_map<int,std::vector<int>>& graph,std::vector<bool>& visited,int node_num,int height,int visited_num,int target_num,int pre_min){
        if(height > pre_min)
            return height;
        if(visited_num == target_num)
            return height;
        if(visited[node_num])
            return height;
        visited[node_num] = true;
        visited_num++;
        int max_height = height;
        for(int i = 0;i < graph[node_num].size();i++){
            if(visited[graph[node_num][i]])
                continue;
            int h = dfs_findMinHeightTrees(graph,visited,graph[node_num][i],height + 1,visited_num,target_num,pre_min);
            if(h > max_height)
                max_height = h;
        }
        return max_height;
    }

    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        std::vector<int> res;
        std::unordered_map<int,std::vector<int>> graph;
        for(int i = 0;i < edges.size();i++){
            graph[edges[i][0]].push_back(edges[i][1]);
            graph[edges[i][1]].push_back(edges[i][0]);
        }
        int min_height = n;
        std::unordered_map<int,std::vector<int>> record;
        std::vector<bool> visited(n,false);
        for(int i = 0;i < n;i++){
            int height = dfs_findMinHeightTrees(graph,visited,i,0,0,n,min_height);
            if(height < min_height)
                min_height = height;
            record[height].push_back(i);
            visited.assign(n,false);
        }
        return record[min_height];
    }
};