#include <vector>
#include <unordered_map>

//The rabbits have same color who tells same numbers.
class Solution_781 {
public:
    int numRabbits(vector<int>& answers){
        std::unordered_map<int,int> record;
        for(auto a : answers){
            record[a]++;
        }
        int cnt = 0;
        for(auto r : record){
            if(r.second <= r.first + 1){
                cnt += r.first + 1;
            }else{
                int k = r.second / (r.first + 1);
                if (r.second % (r.first + 1) != 0){
                    k++;
                }
                cnt += k * (r.first + 1);
//                int total_cnt = r.second;
//                while(total_cnt > 0){
//                    cnt += r.first + 1;
//                    total_cnt -= (r.first + 1);
//                }
            }
        }
        return cnt;
    }
    
//    int numRabbits(vector<int>& answers) {
//        std::unordered_map<int,int> record;
//        for(auto a : answers){
//            record[a]++;
//        }
//        int rabbit_cnt = 0;
//        for(auto r : record){
//            //record[4] = 3
//            if(r.first >= r.second + 1){// They have same color
//                rabbit_cnt += r.first + 1;
//            }else{
//                int total_cnt = r.second;
//                while(total_cnt > 0){
//                    rabbit_cnt += r.first + 1;
//                    total_cnt -= (r.first + 1);
//                }
//            }
//        }
//        return rabbit_cnt;
//    }
};