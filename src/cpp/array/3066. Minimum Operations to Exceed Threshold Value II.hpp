#include <vector>
#include <queue>
using namespace std;

//你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。
//nums = [2,11,10,1,3], k = 10
class Solution_3066 {
public:
    int minOperations(vector<int>& nums, int k) {
        std::priority_queue<long,vector<long>,greater<long>> q;//small top stack
        for(auto n : nums){
            q.push(n);
        }
        int cnt = 0;
        while(q.size() >= 2 && q.top() < k){
            auto n1 = q.top();
            q.pop();
            auto n2 = q.top();
            q.pop();
            long n3 = min(n1, n2) * 2 + max(n1, n2);
            q.push(n3);
            cnt++;
        }
        return cnt;
    }
};