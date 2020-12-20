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

    //greedy solution is wrong
//    int lastStoneWeightII(vector<int>& stones) {
//        std::priority_queue<int> q{stones.begin(),stones.end()};
//        while(q.size() > 1){
//            auto big1 = q.top();
//            q.pop();
//            auto big2 = q.top();
//            q.pop();
//            if (big1 != big2){
//                q.push(big1 - big2);
//            }
//        }
//        if(q.size() == 0)
//            return 0;
//        return q.top();
//    }

//dfs solution will TLE
//    int dp_lastStoneWeightII(vector<int>& stones){
//        int len = stones.size();
//        if (len == 0)
//            return 0;
//        if (len == 1)
//            return stones[0];
//        int res = 2147483647;
//        for(int i = 0;i < len;i++){
//            for(int j = i + 1;j < len;j++){
//                int first = stones[i];
//                int second = stones[j];
//                std::vector<int> new_stones;
//                int index = 0;
//                int len2 = 0;
//                if (first == second){
//                    len2 = len - 2;
//                }else{
//                    len2 = len - 1;
//                }
//                new_stones.resize(len2);
//                for(int k = 0;k < len;k++){
//                    if(k != i && k != j)
//                        new_stones[index++] = stones[k];
//                }
//                if(first != second)
//                    new_stones[index] = abs(first - second);
//                int cur = dp_lastStoneWeightII(new_stones);
//                if(cur < res)
//                    res = cur;
//            }
//        }
//        return res;
//    }
//
//    int lastStoneWeightII(vector<int>& stones){
//        return dp_lastStoneWeightII(stones);
//    }
};