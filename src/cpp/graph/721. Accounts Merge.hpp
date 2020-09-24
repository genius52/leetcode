#include <vector>
#include <string>
#include <unordered_map>
#include <set>
#include <sstream>
using namespace std;

//Input:
//accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
//Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
class Solution_721 {
public:
    vector<vector<string>> accountsMerge(vector<vector<string>>& accounts) {
        int len = accounts.size();
        std::vector<std::vector<string>> res;
        if (len == 0){
            return res;
        }
        std::unordered_map<string,std::set<string>> people_2_mails;//name : mails
        std::unordered_map<string,std::set<int>> mail_2_people;//mail -> people index
        std::unordered_map<int,std::set<int>> relation;
        std::vector<bool> visited;
        visited.resize(len);
        for (int i = 0;i < len;i++){
            for(int j = 1;j < accounts[i].size();j++){
                if(mail_2_people.find(accounts[i][j]) != mail_2_people.end()){
                    bool dup = false;
                    for(auto index : mail_2_people[accounts[i][j]]){
                        if (i == index){
                            dup = true;
                            break;
                        }
                    }
                    if(!dup){
                        for (auto index : mail_2_people[accounts[i][j]]){
                            relation[i].insert(index);
                            relation[index].insert(i);
                        }
                    }
                }
                mail_2_people[accounts[i][j]].insert(i);
            }
        }
        //for(auto it = relation.begin();it != relation.end();it++)
        for(int i = 0;i < len;i++){
            if(relation.find(i) != relation.end()){
                if(!visited[i]){
                    int l = accounts[i].size();
                    for(int j = 1;j < l;j++){
                        people_2_mails[accounts[i][0]].insert(accounts[i][j]);
                    }
                    visited[i] = true;
                }
                std::set<int> same = relation[i];
                for(auto s:same){
                    if(visited[s]){
                        continue;
                    }
                    int l = accounts[s].size();
                    for(int j = 1;j < l;j++){
                        people_2_mails[accounts[s][0]].insert(accounts[s][j]);
                    }
                    visited[s] = true;
                }
            }
        }
        for(auto p:people_2_mails){
            std::vector<std::string> v;
            v.push_back(p.first);
            for(auto m : p.second){
                v.push_back(m);
            }
            res.push_back(v);
        }
        for(int i = 0;i < len;i++){
            if(visited[i])
                continue;
            std::vector<std::string> v;
            v.push_back(accounts[i][0]);
            std::set<std::string> s;
            int l = accounts[i].size();
            for(int j = 1;j < l;j++){
                s.insert(accounts[i][j]);
            }
            for(auto e:s){
                v.push_back(e);
            }
            res.push_back(v);
        }
        return res;
    }
};