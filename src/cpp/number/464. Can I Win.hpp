#include <unordered_map>
#include <set>
using namespace std;
//Input:
//maxChoosableInteger = 10
//desiredTotal = 11
//
//Output:
//false
//
//Explanation:
//No matter which integer the first player choose, the first player will lose.
//The first player can choose an integer from 1 up to 10.
//If the first player choose 1, the second player can only choose integers from 2 up to 10.
//The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
//Same with other integers chosen by the first player, the second player will always win.
class Solution_464 {
public:
    bool dfs_canIWin(int desiredTotal,bool isme,std::set<int> s,std::unordered_map<int,bool>& record){
        if(record.find(desiredTotal) != record.end())
            return record[desiredTotal];
        if(s.empty()){
            if(isme){
                record[desiredTotal] = false;
                return false;
            }
            else{
                record[desiredTotal] = true;
                return true;
            }
        }
        if(*s.rbegin() >= desiredTotal){
            if(isme){
                record[desiredTotal] = true;
                return true;
            }else{
                record[desiredTotal] = false;
                return false;
            }
        }
        bool res = false;
        for(auto it = s.begin();it != s.end();it++){
            int val = *it;
            s.erase(val);
            res |= dfs_canIWin(desiredTotal - val,!isme,s,record);
            if(res){
                break;
            }
            s.insert(val);
        }
        record[desiredTotal] = res;
        return res;
    }

    bool canIWin(int maxChoosableInteger, int desiredTotal) {
        std::unordered_map<int,bool> record;
        std::set<int> s;
        for(int i = 1;i < maxChoosableInteger;i++){
            s.insert(i);
        }
        return dfs_canIWin(desiredTotal,true,s,record);
    }
};