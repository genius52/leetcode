#include <vector>
#include <map>

//The rabbits have same color who tells same numbers.
class Solution_781 {
public:
    int numRabbits(vector<int>& answers) {
        std::map<int,int> record;
        for(auto a : answers){
            record[a]++;
        }
        int rabbit_cnt = 0;
        for(auto r : record){
            if(r.first >= r.second + 1){
                rabbit_cnt += r.first + 1;
            }else{
                int total = r.second;
                while(total > 0){
                    rabbit_cnt += r.first + 1;
                    total -= (r.first + 1);
                }
            }
        }
        return rabbit_cnt;
    }
};