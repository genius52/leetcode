#include <vector>
#include <queue>
using namespace std;


class Solution_1049 {
public:
    //1 <= stones.length <= 30
    //1 <= stones[i] <= 100
    int lastStoneWeightII(vector<int>& stones){
        std::vector<bool> dp;
        dp.resize(1501);//dp[num] means achieve num is available
        int len = stones.size();
        int sum = 0;
        for(int i = 0;i < len;i++){
            sum += stones[i];
            dp[sum] = true;
        }
        int target = sum / 2;
        for(int i = 0;i < len;i++){

        }
    }
};