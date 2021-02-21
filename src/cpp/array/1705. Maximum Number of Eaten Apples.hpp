#include <vector>
#include <map>
using namespace std;

//apples = [1,2,3,5,2], days = [3,2,1,4,2]
class Solution_1705 {
public:
    int eatenApples(vector<int>& apples, vector<int>& days) {
        std::map<int,int> record;//key: date, val:count
        int len = apples.size();
        int total = 0;
        int res = 0;
        for(int i = 0;i < len;i++){
            if(apples[i] != 0){
                record[i + days[i] - 1] += apples[i];
                total += apples[i];
            }
            if(total > 0){
                auto it = record.lower_bound(i);
                if(it->first == i || it->second == 1){
                    total -= it->second;
                    res++;
                    record.erase(it);
                }else{
                    res++;
                    it->second--;
                    if(it->second == 0){
                        record.erase(it);
                    }
                    total--;
                }
            }
        }
        int date = len - 1;
        for(auto it = record.begin();it != record.end();it++){
            int m = min(it->first - date,it->second);
            res += m;
            if(it->second >= (it->first - date)){
                date = it->first;
            }else{
                date += (it->first - date - it->second);
            }
        }
        return res;
    }
};