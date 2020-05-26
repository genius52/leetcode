#include <vector>
#include <unordered_map>
#include <queue>
//Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
//Output: false
class Solution_207 {
public:
    bool dfs_canFinish(int cur_course,std::unordered_map<int,std::vector<int>>& record,std::vector<bool>& trace,std::vector<bool>& visited){
        if (trace[cur_course])
            return false;
        if(visited[cur_course])
            return true;
        trace[cur_course] = true;
        if(record.find(cur_course) != record.end())
        {
            for(auto it = record[cur_course].begin();it != record[cur_course].end();it++){
                if(!dfs_canFinish(*it,record,trace,visited))
                    return false;
            }
        }
        trace[cur_course] = false;
        visited[cur_course] = true;
        return true;
    }

    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        int len = prerequisites.size();
        std::unordered_map<int,std::vector<int>> record;
        for(int i = 0;i < len;i++){
            if(record.find(prerequisites[i][0]) != record.end()){
                record[prerequisites[i][0]].push_back(prerequisites[i][1]);
            }
            else{
                record[prerequisites[i][0]] = std::vector<int>{prerequisites[i][1]};
            }
        }
        std::vector<bool> visited;
        visited.resize(numCourses);
        for(int i = 0;i < numCourses;i++){
            std::vector<bool> trace;
            trace.resize(numCourses);
            if (!dfs_canFinish(i,record,trace,visited))
                return false;
        }
        return true;
    }

    bool canFinish_bfs(int numCourses, vector<vector<int>>& prerequisites) {
        int len = prerequisites.size();
        std::vector<std::vector<int>> graph;
        graph.resize(numCourses);
        std::vector<int> indegree;//store indegree for each node
        indegree.resize(numCourses);
        for(int i = 0;i < len;i++){
            graph[prerequisites[i][0]].push_back(prerequisites[i][1]);
            indegree[prerequisites[i][1]]++;
        }
        std::queue<int> q;
        for(int i = 0;i < numCourses;i++){
            if(indegree[i] == 0){
                q.push(i);
            }
        }
        while(!q.empty()) {
            auto n = q.front();
            {
                q.pop();
                for (int i = 0; i < graph[n].size(); i++) {
                    indegree[graph[n][i]]--;
                    if (indegree[graph[n][i]] == 0)
                        q.push(graph[n][i]);
                }
            }
        }
        for(int i = 0;i < numCourses;i++){
            if(indegree[i] > 0)
                return false;
        }
        return true;
    }
};