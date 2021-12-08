#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

class Solution_1751 {
public:
    int dp_maxValue(vector<vector<int>>& events,int len,int pos,int last,int k,std::unordered_map<int,std::unordered_map<int,int>>& memo){
        if(pos >= len)
            return 0;
        if (k == 0)
            return 0;
        if(events[pos][0] <= last){
            return dp_maxValue(events,len,pos + 1,last,k,memo);
        }else{
            if(memo.find(pos) != memo.end()){
                if(memo[pos].find(k) != memo[pos].end()){
                    return memo[pos][k];
                }
            }
            //choose current
            int res1 = events[pos][2] + dp_maxValue(events,len,pos + 1,events[pos][1],k - 1,memo);

            //skip current
            int res2 = dp_maxValue(events,len,pos + 1,last,k,memo);
            memo[pos][k] = max(res1,res2);
            return memo[pos][k];
        }
    }

    int maxValue(vector<vector<int>>& events, int k) {
        std::sort(events.begin(),events.end());
        std::unordered_map<int,std::unordered_map<int,int>> memo;
        return dp_maxValue(events,events.size(),0,-1,k,memo);
    }
};