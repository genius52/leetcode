#include <vector>
#include <map>
using namespace std;

class Solution_1887 {
public:
    int reductionOperations(vector<int>& nums) {
        std::map<int,int> record;
        int min_num = 2147483647;
        for(auto n : nums){
            record[n]++;
            if(n < min_num)
                min_num = n;
        }
        int steps = 0;
        int total_cnt = 0;
        auto it = record.rbegin();
        while((*it).first != min_num){
            int cur_cnt = (*it).second;
            total_cnt += cur_cnt;
            steps += total_cnt;
            it++;
        }
        return steps;
    }
};