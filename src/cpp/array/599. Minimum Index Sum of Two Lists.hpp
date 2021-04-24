#include <vector>
#include <string>
#include <unordered_set>
#include <unordered_map>
#include <map>
using namespace std;

class Solution_599 {
public:
    vector<string> findRestaurant(vector<string>& list1, vector<string>& list2) {
        int len1 = list1.size();
        int len2 = list2.size();
        std::unordered_map<std::string,int> record1;
        std::unordered_map<std::string,int> record2;
        for(int i = 0;i < len1;i++){
            record1[list1[i]] = i;
        }
        for(int i = 0;i < len2;i++){
            record2[list2[i]] = i;
        }
        std::map<int,std::vector<std::string>> index_string;
        for(auto it = record1.begin();it != record1.end();it++){
            if(record2.find(it->first) == record2.end())
                continue;
            index_string[it->second + record2[it->first]].push_back(it->first);
        }
        return index_string.begin()->second;
    }
};