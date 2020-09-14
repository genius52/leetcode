#include <vector>
#include <map>
using namespace std;

class Solution_659 {
public:
    bool isPossible(vector<int>& nums) {
        std::map<int,int> record;
        for(auto n : nums){
            record[n]++;
        }
        while(!record.empty()){
            int len = 0;
            int last_dup_cnt = 0;
            int last_num = 2147483647;
            for(auto it = record.begin();it != record.end();){
                if (last_num == 2147483647){
                    last_num = it->first;
                }
                else{
                    if(it->first - last_num != 1){
                        break;
                    }
                    last_num = it->first;
                }
                if (it->second >= last_dup_cnt){
                    len++;
                    last_dup_cnt = it->second;
                    it->second--;
                    if(it->second == 0){
                        record.erase(it++);
                    }else{
                        it++;
                    }
                }else{
                    break;
                }
            }
            if (len < 3){
                return false;
            }
        }
        return true;
    }
};