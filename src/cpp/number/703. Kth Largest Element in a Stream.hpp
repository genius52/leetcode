#include <vector>
#include <queue>
using namespace std;

class KthLargest {
public:
    KthLargest(int k, vector<int>& nums) {
        kth = k;
        int len = nums.size();
        for(auto n : nums){
            if(q.size() < k){
                q.push(n);
            }else{
                if(n > q.top()){
                    q.pop();
                    q.push(n);
                }
            }
        }
    }
    int add(int val) {
        if(q.size() < kth){
            q.push(val);
            return q.top();
        }
        if(val >= q.top()){
            q.pop();
            q.push(val);
        }
        return q.top();
    }
    int kth = 0;
    std::priority_queue<int, vector<int>, greater<int> > q;
};