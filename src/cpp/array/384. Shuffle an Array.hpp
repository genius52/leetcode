#include <vector>
#include <math.h>
using namespace std;

class Solution_384 {
    std::vector<int> data;
public:
    Solution(vector<int>& nums) {
        data = nums;
    }

    /** Resets the array to its original configuration and return it. */
    vector<int> reset() {
        return data;
    }

//第1次从1-n的候选集合随机选个数，拿出此数，并把它从候选集合剔除(候选集合n-1)。
//第2次从1-n-1的候选集合随机选个数，拿出此数，并把它从候选集合剔除(候选集合n-2)。
//第2次从1-n-2的候选集合随机选个数，拿出此数，并把它从候选集合剔除(候选集合n-3)。
    /** Returns a random shuffling of the array. */
    vector<int> shuffle() {
        int len = data.size();
        auto res = data;
        for(int i = 0;i < len;i++){
            int swap_pos = rand() % (len - i)+ i;
            auto tmp = res[i];
            res[i] = res[swap_pos];
            res[swap_pos] = tmp;
        }
        return res;
    }
};