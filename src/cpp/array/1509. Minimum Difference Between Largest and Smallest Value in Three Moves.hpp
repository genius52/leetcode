#include <vector>
#include <map>
using namespace std;

//Input: nums = [5,3,2,4]
//Output: 0
//Explanation: Change the array [5,3,2,4] to [2,2,2,2].
//The difference between the maximum and minimum is 2-2 = 0.
class Solution_1509 {
public:
    int dp_minDifference(int steps,std::map<int,int>::iterator it1,std::map<int,int>::reverse_iterator it2,std::map<int,int> record){
        if (steps == 0){
            return abs((*it1).first - (*it2).first);
        }
        if (it1 == record.end() || it2 == record.rend()){
            return 0;
        }
        if ((*it1).first == (*it2).first){
            return 0;
        }
        int res = 2147483647;
        int cnt1 = (*it1).second;
        int cnt2 = (*it2).second;
        if (cnt1 > 1){
            (*it1).second--;
            res = min(res,dp_minDifference(steps - 1,it1,it2,record));
            (*it1).second++;
        }else{
            auto new_it1 = it1;
            new_it1++;
            res = min(res,dp_minDifference(steps - 1,new_it1,it2,record));
        }
        if(cnt2 > 1){
            (*it2).second--;
            res = min(res,dp_minDifference(steps - 1,it1,it2,record));
            (*it2).second++;
        }else{
            auto new_it2 = it2;
            new_it2++;
            res = min(res,dp_minDifference(steps - 1,it1,new_it2,record));
        }
        return res;
    }

    int minDifference(vector<int>& nums) {
        int len = nums.size();
        if (len <= 4){
            return 0;
        }
        std::map<int,int> record;
        for(auto n : nums){
            record[n]++;
        }
        auto it1 = record.begin();
        auto it2 = record.rbegin();
        return dp_minDifference(3,record.begin(),record.rbegin(),record);
    }
};