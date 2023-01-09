#include <vector>
#include <queue>
#include <math.h>
using namespace std;

class Solution_2530 {
public:
    long long maxKelements(vector<int>& nums, int k) {
        std::priority_queue<int,vector<int>,std::less<int>> q(nums.begin(),nums.end());
        long long res = 0;
        while (k > 0){
            int val = q.top();
            res += val;
            q.pop();
            q.push(ceil(double(val)/3));
            k--;
        }
        return res;
    }
};