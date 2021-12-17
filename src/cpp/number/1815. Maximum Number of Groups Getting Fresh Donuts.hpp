#include <vector>
#include <map>
using namespace std;

class Solution_1815 {
public:
    int dp_maxHappyGroups(std::vector<int>& record,int batchSize,int rest,std::map<std::vector<int>,int>& memo){
        if(memo.find(record) != memo.end())
            return memo[record];
        int res = 0;
        for(int i = 1;i < record.size();i++){
            if(record[i] == 0){
                continue;
            }
            record[i]--;
            int cur = 0;
            if(rest == 0){
                cur = 1;
            }
            res = max(res,cur + dp_maxHappyGroups(record,batchSize,(rest + i) % batchSize,memo));
            record[i]++;
        }
        memo[record] = res;
        return res;
    }

    int maxHappyGroups(int batchSize, vector<int>& groups) {
        std::vector<int> record(batchSize);
        int res = 0;
        for(auto n : groups){
            int mod = n % batchSize;
            if(mod == 0){
                res++;
                continue;
            }else if(record[batchSize - mod] != 0){
                record[batchSize - mod]--;
                res++;
            }else{
                record[mod]++;
            }
        }
        std::map<std::vector<int>,int> memo;
        return res + dp_maxHappyGroups(record,batchSize,0,memo);
    }
};