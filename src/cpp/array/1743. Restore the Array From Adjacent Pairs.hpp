#include <vector>
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution_1743 {
public:
    vector<int> restoreArray(vector<vector<int>>& adjacentPairs) {
        int len = adjacentPairs.size();
        std::unordered_map<int,int> count;
        std::unordered_map<int,std::vector<int>> graph;
        for(int i = 0;i < len;i++){
            if(graph.find(adjacentPairs[i][0]) == graph.end()){
                graph[adjacentPairs[i][0]] = std::vector<int>{adjacentPairs[i][1]};
            }else{
                graph[adjacentPairs[i][0]].push_back(adjacentPairs[i][1]);
            }
            if(graph.find(adjacentPairs[i][1]) == graph.end()){
                graph[adjacentPairs[i][1]] = std::vector<int>{adjacentPairs[i][0]};
            }else{
                graph[adjacentPairs[i][1]].push_back(adjacentPairs[i][0]);
            }
            count[adjacentPairs[i][0]]++;
            count[adjacentPairs[i][1]]++;
        }
        std::vector<int> res(len + 1);
        int single[2] = {0};
        int idx = 0;
        for(auto it = count.begin();it != count.end();it++){
            if((*it).second == 1){
                single[idx++] = (*it).first;
                if(idx == 2)
                    break;
            }
        }
        res[0] = single[0];
        res[len] = single[1];
        int start = res[0];
        int end = res[len];
        int pos = 1;
        std::unordered_set<int> visited;
        visited.insert(start);
        visited.insert(end);
        while(pos < len){
            std::vector<int> neighbours = graph[start];
            for(auto it = neighbours.begin();it != neighbours.end();it++){
                if(visited.find(*it) == visited.end()){
                    start = *it;
                    res[pos] = *it;
                    visited.insert(*it);
                    pos++;
                    break;
                }
            }
        }
        return res;
    }
};