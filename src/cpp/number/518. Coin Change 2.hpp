#include <vector>
using namespace std;

class Solution_518 {
public:
    int dfs_change(int amount,std::vector<int>& coins,int len,int pos,std::unordered_map<int,std::unordered_map<int,int>>& memo){
        if(amount == 0){
            return 1;
        }
        if(pos >= len){
            return 0;
        }
        if(memo.find(amount) != memo.end()){
            if(memo[amount].find(pos) != memo[amount].end()){
                return memo[amount][pos];
            }
        }
        int total = 0;
        if(coins[pos] > amount){
            total = dfs_change(amount,coins,len,pos + 1,memo);
        }else{
            for(int i = 0;amount - i * coins[pos] >= 0;i++){
                int rest = amount - i * coins[pos];
                if(rest < 0)
                    break;
                total += dfs_change(rest,coins,len,pos + 1,memo);
            }
        }
        memo[amount][pos] = total;
        return total;
    }

    int change(int amount, vector<int>& coins) {
        int len = coins.size();
        std::unordered_map<int,std::unordered_map<int,int>> memo;
        return dfs_change(amount,coins,len,0,memo);
    }
};