#include <vector>
#include <string>
#include <map>
#include <set>
using namespace std;

class Solution_692 {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
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