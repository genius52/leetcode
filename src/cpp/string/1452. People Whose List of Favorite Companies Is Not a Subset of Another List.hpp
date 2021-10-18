#include <vector>
#include <string>
#include <unordered_map>
#include <bitset>
using namespace std;

class Solution_1452 {
public:
    vector<int> peopleIndexes(vector<vector<string>>& favoriteCompanies) {
        std::unordered_map<std::string,int> name_idx;
        int len = favoriteCompanies.size();
        std::vector<std::bitset<50001>> record(len);
        int idx = 1;
        for(int i = 0;i < len;i++){
            std::bitset<50001> cur;
            for(int j = 0;j < favoriteCompanies[i].size();j++){
                if(name_idx.find(favoriteCompanies[i][j]) == name_idx.end()){
                    name_idx[favoriteCompanies[i][j]] = idx++;
                }
                cur.set(name_idx[favoriteCompanies[i][j]],1);
            }
            record[i] = cur;
        }
        std::vector<int> res;
        for(int i = 0;i < len;i++){
            bool match = true;
            for(int j = 0;j < len;j++){
                if(i == j)
                    continue;
                if((record[i] | record[j]) == record[j]){
                    match = false;
                    break;
                }
            }
            if(match)
                res.push_back(i);
        }
        return res;
    }
};