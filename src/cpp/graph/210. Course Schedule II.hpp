#include <vector>
#include <unordered_map>
#include <deque>
//using namespace std;

//Input: 4, [[1,0],[2,0],[3,1],[3,2]]
//Output: [0,1,2,3] or [0,2,1,3]
class Solution_210 {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        std::unordered_map<int,std::vector<int>> graph;
        std::vector<int> degree;//out degree?
        degree.resize(numCourses);
        for(int i = 0;i < prerequisites.size();i++){
            degree[prerequisites[i][0]]++;
            graph[prerequisites[i][1]].push_back(prerequisites[i][0]);
        }
        std::vector<int> res;
        std::deque<int> q;
        for(int i = 0;i < numCourses;i++){
            if(degree[i] == 0){
                q.push_back(i);
                res.push_back(i);
            }
        }
        while(!q.empty()){
            int n = q.front();
            q.pop_front();
            if(graph.find(n) != graph.end()){
                for(int i = 0;i < graph[n].size();i++){
                    degree[graph[n][i]]--;
                    if(degree[graph[n][i]] == 0){
                        q.push_back(graph[n][i]);
                        res.push_back(graph[n][i]);
                    }
                }
            }
        }
        for(int i = 0;i < numCourses;i++){
            if(degree[i] != 0){
                return vector<int>{};
            }
        }
        return res;
    }
};