#include <string>
#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution_433 {
public:
    bool is_mutation(std::string& s1,std::string& s2){
        int len1 = s1.size();
        int len2 = s2.size();
        if(len1 != len2)
            return false;
        int diff_cnt = 0;
        for(int i = 0;i < len1;i++){
            if(s1[i] != s2[i])
            {
                diff_cnt++;
                if(diff_cnt > 1)
                    return false;
            }
        }
        return true;
    }
    bool dfs_minMutation(string start,string end,int& total,std::unordered_map<std::string,std::unordered_set<std::string>>& graph,std::unordered_set<std::string>& visited){
        if(start == end)
            return true;
        int min_step = 1<<30;
        for(auto it = graph[start].begin();it != graph[start].end();it++){
            int step = 0;
            auto next = *it;
            if(visited.find(next) != visited.end())
                continue;
            visited.insert(next);
            bool ret = dfs_minMutation(next,end,step,graph,visited);
            visited.erase(next);
            if(ret){
                if(step + 1 < min_step)
                    min_step = step + 1;
            }
        }
        if(min_step == 1<<30)
            return false;
        total = min_step;
        return true;
    }

    int minMutation(string start, string end, vector<string>& bank) {
        int l = bank.size();
        std::unordered_map<std::string,std::unordered_set<std::string>> graph;
        for(int i = 0;i < l;i++){
            for(int j = 0;j < l;j++){
                if(i == j)
                    continue;
                if(is_mutation(bank[i],bank[j]))
                    graph[bank[i]].insert(bank[j]);
            }
        }
        for(int i = 0;i < l;i++){
            if(is_mutation(bank[i],start))
                graph[start].insert(bank[i]);
        }
        int res = 0;
        std::unordered_set<std::string> visited;
        if(dfs_minMutation(start,end,res,graph,visited))
            return res;
        return -1;
    }

    //int minMutation2(string start, string end, vector<string>& bank) {
    //    int l = bank.size();
    //    std::unordered_map<std::string,std::unordered_set<std::string>> graph;
    //    for(int i = 0;i < l;i++){
    //        for(int j = 0;j < l;j++){
    //            if(i == j)
    //                continue;
    //            if(is_mutation(bank[i],bank[j]))
    //                graph[bank[i]].insert(bank[j]);
    //        }
    //    }
    //    for(int i = 0;i < l;i++){
    //        if(is_mutation(bank[i],start))
    //            graph[start].insert(bank[i]);
    //    }
    //    int res = 0;
    //    std::unordered_set<std::string> visited;
    //    std::queue<std::string> q;
    //    q.push(start);
    //    int cnt = 0;
    //    while(q.empty()){
    //        q.front();
    //    }
    //    return cnt;
    //}
};