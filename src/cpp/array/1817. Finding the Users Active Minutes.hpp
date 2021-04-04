#include <vector>
#include <unordered_set>
#include <map>
#include <string>
using namespace std;

class Solution_1817 {
public:
    vector<int> findingUsersActiveMinutes(vector<vector<int>>& logs, int k) {
        std::unordered_map<int,std::unordered_set<int>> record;
        int len = logs.size();
        for(int i = 0;i < len;i++){
            record[logs[i][0]].insert(logs[i][1]);
        }
        std::map<int,int> minute_count;
        for(auto it = record.begin();it != record.end();it++){
            int minutes = it->second.size();
            minute_count[minutes]++;
        }
        std::vector<int> res(k);
        for(int i = 1;i <= k;i++){
            if(minute_count.find(i) != minute_count.end()){
                res[i - 1] = minute_count[i];
            }else{
                res[i - 1] = 0;
            }
        }
        return res;
    }
};