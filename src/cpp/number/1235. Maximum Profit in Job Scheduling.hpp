#include <vector>
#include <unordered_map>
using namespace std;

class Solution_1235 {
public:
    int dp_jobScheduling(std::vector<std::vector<int>>& record,int len,int cur_pos,std::unordered_map<int,int>& memo){
        if(cur_pos >= len)
            return 0;
        //choose next task
        if(memo.find(cur_pos) != memo.end())
            return memo[cur_pos];
        int idx = cur_pos + 1;
        while(idx < len){
            if(record[idx][0] >= record[cur_pos][1])
                break;
            idx++;
        }
        int res1 = record[cur_pos][2] + dp_jobScheduling(record,len,idx,memo);
        //skip cur task
        int res2 = dp_jobScheduling(record,len,cur_pos + 1,memo);
        memo[cur_pos] = max(res1,res2);
        return memo[cur_pos];
    }
    int jobScheduling(vector<int>& startTime, vector<int>& endTime, vector<int>& profit) {
        int len = profit.size();
        std::vector<std::vector<int>> record(len);
        for(int i = 0;i < len;i++){
            record[i] = std::vector<int>{startTime[i],endTime[i],profit[i]};
        }
        std::sort(record.begin(),record.end(),[](std::vector<int>& v1,std::vector<int>& v2){
            if(v1[0] == v2[0]){
                return v1[1] < v2[1];
            }
            return v1[0] < v2[0];
        });
        std::unordered_map<int,int> memo;
        return dp_jobScheduling(record,len,0,memo);
    }
};