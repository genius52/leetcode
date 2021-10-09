#include <vector>
#include <string>
#include <map>
#include <unordered_map>
#include <set>
#include <list>
using namespace std;

class WordFilter {
    std::unordered_map<std::string,std::set<int>> prefix_map;
    std::unordered_map<std::string,std::set<int>> suffix_map;
public:
    WordFilter(vector<string>& words) {
        int len = words.size();
        for(int i = 0;i < len;i++){
            int word_len = words[i].length();
            for(int j = 1;j <= word_len;j++){
                prefix_map[words[i].substr(0,j)].insert(i);
                suffix_map[words[i].substr(word_len - j,j)].insert(i);
            }
        }
    }

    int f(string prefix, string suffix) {
        auto it1 = prefix_map.find(prefix);
        auto it2 = suffix_map.find(suffix);
        if(it1 == prefix_map.end() || it2 == suffix_map.end()){
            return -1;
        }
        for(auto it11 = prefix_map[prefix].rbegin();it11 != prefix_map[prefix].rend();it11++){
            if(suffix_map[suffix].find(*it11) != suffix_map[suffix].end())
                return *it11;
        }
        return -1;
    }
};