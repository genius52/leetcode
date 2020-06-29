#include <math.h>
//Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
//Output: true
//Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
class Solution_1497 {
public:
    bool canArrange(vector<int>& arr, int k) {
        std::map<int,int> record;
        int len = arr.size();
        int r = -81199994 % 100000;
        for(auto n : arr){
            auto rest = n % k;
            record[(rest + k) % k]++;
        }
        for(auto it = record.begin();it != record.end();it++){
            if((*it).second == 0)
                continue;
            auto num = (*it).first;
            auto cnt = (*it).second;
            if(num == 0){
                if(cnt % 2 != 0)
                    return false;
                continue;
            }
            int need = k - num;
            if(record.find(need) == record.end())
                return false;
            if(record[need] < cnt)
                return false;
            record[need] -= cnt;
        }
        return true;
    }
};