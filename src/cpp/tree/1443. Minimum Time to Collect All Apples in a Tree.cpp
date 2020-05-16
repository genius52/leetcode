//
// Created by 陶诚程 on 2020-05-16.
//
#include <vector>
using namespace std;
//Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
//Output: 8
class Solution_1443 {
public:
    int dfs_minTime(int node,int step,vector<vector<int>>& relation,vector<bool>& hasApple,std::vector<bool>& visited){
        if (visited[node])
            return 0;
        visited[node] = true;
        auto children = relation[node];
        int children_steps = 0;
        for(int i = 0;i < children.size();i++){
            if(visited[children[i]]){
                continue;
            }
            children_steps += dfs_minTime(children[i],2,relation,hasApple,visited);
        }
        if(!hasApple[node] && children_steps == 0){
            return 0;
        }
        return step + children_steps;
    }

    int minTime(int n, vector<vector<int>>& edges, vector<bool>& hasApple) {
        std::vector<std::vector<int> > adjacent;//(n,std::vector<bool>(n,false));
        adjacent.resize(n);
        for (auto r : edges){
            adjacent[r[0]].push_back(r[1]);
            adjacent[r[1]].push_back(r[0]);
        }
        std::vector<bool> visited;
        visited.resize(n);
        return dfs_minTime(0,0,adjacent,hasApple,visited);
    }
};