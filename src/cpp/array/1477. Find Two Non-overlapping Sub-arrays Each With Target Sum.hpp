#include <vector>
#include <unordered_map>
using namespace std;

//Input: arr = [3,1,1,1,5,1,2,1], target = 3
//Output: 3
class Solution_1477 {
public:
    int minSumOfLengths(vector<int>& arr, int target) {
        int len = arr.size();
        if(len < 1)
            return -1;
        std::vector<int> dp(len + 1);
        std::unordered_map<int,int> record;
        dp[0] = arr[i];
        record[dp[0]] = 0;
        for(int i = 1;i < len;i++){
            dp[i] = dp[i - 1] + arr[i];
            record[dp[i + 1]] = i;
        }
//        std::set<int> s;
//        for(int i = 0;i < len - 1;i++){
//            for(int j = i + 1;j < len;j++){
//
//            }
//        }
    }
};