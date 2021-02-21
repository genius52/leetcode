#include <vector>
#include <queue>
using namespace std;

class Solution_1760 {
public:
    int minimumSize(vector<int>& nums, int maxOperations) {
        std::priority_queue<int> q(nums.begin(),nums.end());
        while (maxOperations > 0){
            int top = q.top();
            q.pop();
            int num1 = top/2;
            q.push(num1);
            if((top|1) == top){
                q.push(num1 + 1);
            }else{
                q.push(num1);
            }
        }
        return q.top();
    }
};