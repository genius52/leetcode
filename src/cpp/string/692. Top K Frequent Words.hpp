#include <vector>
#include <string>
#include <map>
#include <set>
#include <queue>
using namespace std;

class Solution_692 {
public:
    vector<string> topKFrequent(vector<string>& words, int k){
        std::unordered_map<std::string,int> word_cnt;
        for(auto word : words){
            word_cnt[word]++;
        }
        std::vector<string> res(k);
        auto my_cmp = [](const std::pair<int,std::string>& p1,const std::pair<int,std::string>& p2)->bool{
            if(p1.first == p2.first){
                return p1.second > p2.second;
            }else{
                return p1.first < p2.first;
            }
        };
        std::priority_queue<std::pair<int,std::string>,std::vector<std::pair<int,std::string>>, decltype(my_cmp)>  q(my_cmp);//weight : index
        for(auto e : word_cnt){
            q.push({e.second,e.first});
        }
        int i = 0;
        while(k > 0){
            auto cur = q.top();
            res[i] = cur.second;
            q.pop();
            i++;
            k--;
        }
        return res;
    }

    vector<string> topKFrequent2(vector<string>& words, int k) {
        std::unordered_map<std::string,int> data;
        for(auto word : words){
            data[word]++;
        }
        std::map<int,std::set<std::string>> record;
        for(auto it = data.begin();it != data.end();it++){
            record[it->second].insert(it->first);
        }
        std::vector<string> res;
        for(auto it = record.rbegin();it != record.rend() && k > 0;it++){
            for(auto it2 = it->second.begin();it2 != it->second.end() && k > 0;it2++){
                res.push_back(*it2);
                k--;
            }
        }
        return res;
    }
};