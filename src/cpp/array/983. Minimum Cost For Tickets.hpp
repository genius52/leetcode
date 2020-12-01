#include <vector>
#include <set>
#include <unordered_map>
using namespace std;

//Input: days = [1,4,6,7,8,20], costs = [2,7,15]
//Output: 11
//Explanation:
//For example, here is one way to buy passes that lets you travel your travel plan:
//On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
//On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
//On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
//In total you spent $11 and covered all the days of your travel.
class Solution_983 {
public:
    int dp_mincostTickets(std::set<int>& s,std::set<int>::const_iterator it,vector<int>& cost,std::unordered_map<int,int>& memo){
        if (it == s.end())
            return 0;
        auto cur_pos = *it;
        if (memo.find(cur_pos) != memo.end())
            return memo[cur_pos];
        int res = 2147483647;
        for (int i = 0;i < 3;i++){
            int cur_cost = cost[i];
            int next_day = *it;
            if (i == 0){
                next_day++;
            }else if(i == 1){
                next_day += 7;
            }else if(i == 2){
                next_day += 30;
            }
            std::set<int>::iterator next = s.lower_bound(next_day);
            if (next != s.end()){
                cur_cost += dp_mincostTickets(s,next,cost,memo);
            }
            if (cur_cost < res)
                res = cur_cost;
        }
        memo[cur_pos] = res;
        return res;
    }

    int mincostTickets(vector<int>& days, vector<int>& costs) {
        std::set<int> s;
        for(auto d : days){
            s.insert(d);
        }
        std::unordered_map<int,int> memo;
        return dp_mincostTickets(s,s.begin(),costs,memo);
    }
};