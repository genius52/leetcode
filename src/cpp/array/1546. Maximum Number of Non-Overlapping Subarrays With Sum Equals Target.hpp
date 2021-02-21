#include <vector>
#include <map>
#include <unordered_map>
#include <algorithm>
using namespace std;

class Solution_1546 {
public:
    int maxNonOverlapping(vector<int>& nums, int target){
        std::unordered_set<int> record;//sum -> count
        int len = nums.size();
        int cur_sum = 0;
        int cnt = 0;
        for(int i = 0;i < len;i++){
            cur_sum += nums[i];
            if(cur_sum == target || record.find(cur_sum - target) != record.end()){
                cnt++;
                record.clear();
                cur_sum = 0;
            }else{
                record.insert(cur_sum);
            }
        }
        return cnt;
    }

//    int maxNonOverlapping(vector<int>& nums, int target) {
//        int len = nums.size();
//        std::vector<int> prefix(len + 1);
//        for(int i = 1;i <= len;i++){
//            prefix[i] = nums[i - 1] + prefix[i - 1];
//        }
//        std::map<int,int,std::greater<int>> record;//record[i] = before position i,the number of sub array.
//        int max_cnt = 0;
//        for(int i = 1;i <= len;i++){
//            for(int j = i - 1;j >= 0;j--){
//                if(prefix[i] - prefix[j] == target){
//                    auto pre_record = record.lower_bound(j);
//                    if(pre_record != record.end()){
//                        record[i] = pre_record->second + 1;
//                    }else{
//                        record[i] = 1;
//                    }
//
//                    if(record[i] > max_cnt){
//                        max_cnt = record[i];
//                    }
//                    break;
//                }
//            }
//        }
//        return max_cnt;
//    }
};