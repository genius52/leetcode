#include <string>
#include <vector>
using namespace std;

class Solution {
public:
    string destCity(vector<vector<string>>& paths) {
        int len = paths.size();
        std::string res;
        std::map<std::string, std::string> record;
        for (int i = 0; i < len; i++){
            record[paths[i][0]] = paths[i][1];
        }
        std::string start_city = paths[0][0];

        while (record.find(start_city) != record.end()) {
            start_city = record[start_city];
        }
        return start_city;
    }
};