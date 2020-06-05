#include <vector>
#include <unordered_map>
#include <string>
#include <unordered_set>
using namespace std;

class Solution_399 {
public:
    double dfs_calcEquation(std::string& start,std::string& end,double pre_calc,std::unordered_map<string,std::unordered_set<string>>& graph,
                            std::unordered_map<string,double>& calc,std::unordered_set<std::string>& visited){
        if(visited.find(start) != visited.end())
            return -1.0;
        if(graph.find(start) == graph.end())
            return -1.0;
        if(start == end)
            return 1.0;
        visited.insert(start);
        auto s = graph[start];
        for(auto it = s.begin();it != s.end();it++){
            std::string exp = start + "," + *it;
            double val = -1.0;
            if(calc.find(exp) != calc.end()){
                val = calc[exp];
            }
            auto result = pre_calc * val;
            if(*it == end)
                return result;
            string new_start = *it;
            auto res = dfs_calcEquation(new_start,end,result,graph,calc,visited);
            if(res != -1.0)
                return res;
        }
        visited.erase(start);
        return -1.0;
    }

    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        int len = equations.size();
        std::unordered_map<string,std::unordered_set<string>> graph;
        std::unordered_map<string,double> calc;
        for(int i = 0;i < len;i++){
            graph[equations[i][0]].insert(equations[i][1]);
            graph[equations[i][1]].insert(equations[i][0]);
            calc[equations[i][0] + "," + equations[i][1]] = values[i];
            calc[equations[i][1] + "," + equations[i][0]] = double(1 / values[i]);
        }
        vector<double> res;
        for(int i = 0;i < queries.size();i++){
            std::unordered_set<std::string> visited;
            auto ret = dfs_calcEquation(queries[i][0],queries[i][1],1,graph,calc,visited);
            res.push_back(ret);
        }
        return res;
    }
};