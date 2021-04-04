#include <string>
#include <unordered_map>
#include <map>
#include <vector>
using namespace std;

class Solution_451 {
public:
    string frequencySort(string s) {
        std::unordered_map<char,int> count;
        std::map<int,std::vector<char>> record;
        for(auto c : s){
            count[c]++;
        }
        for(auto c: count){
            record[c.second].push_back(c.first);
        }
        std::string res;
        for(auto it = record.rbegin();it != record.rend();it++){
            int cnt = it->first;
            for(auto c : it->second){
                std::string sub(cnt,c);
                res += sub;
            }
        }
        return res;
    }
};