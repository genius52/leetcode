#include <vector>
#include <map>
using namespace std;

class Solution_594 {
public:
    int findLHS(vector<int>& nums) {
        std::map<int,int> record;
        for(auto n : nums){
            record[n]++;
        }
        int max_cnt = 0;
        for(auto it = record.begin();it != record.end();it++){
            auto target = it->first + 1;
            if(record.find(target) != record.end()){
                int cnt = it->second + record[target];
                if(cnt > max_cnt){
                    max_cnt = cnt;
                }
            }
        }
        return max_cnt;
    }
};