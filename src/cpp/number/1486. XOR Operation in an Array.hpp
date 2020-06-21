
//Input: n = 4, start = 3
//Output: 8
//Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
//nums[i] = start + 2*i
class Solution_1486 {
public:
    int xorOperation(int n, int start) {
        int res = 0;
        for(int i = 0;i < n;i++){
            res ^= start + 2 * i;
        }
        return res;
    }
};