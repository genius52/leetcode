#include <vector>
#include <map>
#include <string>
#include <unordered_map>
using namespace std;

class Solution_332 {
public:
    void dfs_findItinerary(std::string& city,std::map<std::string,std::multiset<std::string>>& graph,
                           std::vector<std::string>& trace){
        if(graph.find(city) == graph.end()){
            trace.push_back(city);
            return;
        }

        auto it = graph[city].begin();
        while(!graph[city].empty()){
            auto next = *it;
            graph[city].erase(it++);
            dfs_findItinerary(next,graph,trace);
        }
        trace.push_back(city);
    }

    vector<string> findItinerary(vector<vector<string>>& tickets) {
        std::map<std::string,std::multiset<std::string>> graph;
        int len = tickets.size();
        for(int i = 0;i < len;i++){
            graph[tickets[i][0]].insert(tickets[i][1]);
        }

        std::vector<std::string> trace;
        std::string start = "JFK";
        dfs_findItinerary(start,graph,trace);
        std::reverse(trace.begin(),trace.end());
        return trace;
    }
};