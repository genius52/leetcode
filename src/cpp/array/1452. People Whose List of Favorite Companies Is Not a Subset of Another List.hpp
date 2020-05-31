#include <vector>
#include <unordered_map>
using namespace std;

class Solution_1452 {
public:
    vector<int> peopleIndexes(vector<vector<string>>& favoriteCompanies) {
        std::unordered_map<std::string, int> record;
        int len = favoriteCompanies.size();
        std::vector<int> res;
        for (int i = 0; i < len; i++) {
            for (auto it = favoriteCompanies[i].begin(); it != favoriteCompanies[i].end(); it++) {
                if (record.find(*it) == record.end()) {
                    record[*it] = 1;
                }
                else {
                    record[*it]++;
                }
            }
        }
        for (int i = 0; i < len; i++) {
            bool find = false;
            for (auto it = favoriteCompanies[i].begin(); it != favoriteCompanies[i].end(); it++) {
                if (record[*it] == 1) {
                    find = true;
                    break
                }
            }
            if(find)
                res.push_back(i);
        }
        return res;
    }
};