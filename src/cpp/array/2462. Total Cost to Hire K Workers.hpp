#include <vector>
#include <queue>
#include <set>
using namespace std;

class Solution_2462 {
public:
    long long totalCost(vector<int>& costs, int k, int candidates) {
        int len = costs.size();
        if(k >= len){
            return (long long)std::accumulate(costs.begin(),costs.end(),long(0));
        }
        std::priority_queue<int, std::deque<int>, std::greater<int> > head;
        std::priority_queue<int, std::deque<int>, std::greater<int> > tail;
        int left = 0;
        int right = len - 1;
        while(left <= right && left < candidates){
            if(left == right){
                head.push(costs[left]);
            }else{
                head.push(costs[left]);
                tail.push(costs[right]);
            }
            left++;
            right--;
        }
        long long res = 0;
        while(left <= right && k > 0){
            auto top1 = head.top();
            auto top2 = tail.top();
            if(top1 <= top2){
                res += (long long)(top1);
                head.pop();
                head.push(costs[left]);
                left++;
            }else{
                res += (long long)(top2);
                tail.pop();
                tail.push(costs[right]);
                right--;
            }
            k--;
        }
        while(k > 0){
            if (head.empty()){
                auto top = tail.top();
                res += (long long)top;
                tail.pop();
            }else if(tail.empty()){
                auto top = head.top();
                res += (long long)top;
                head.pop();
            }else{
                auto top1 = head.top();
                auto top2 = tail.top();
                if(top1 <= top2){
                    res += (long long)top1;
                    head.pop();
                }else{
                    res += (long long)top2;
                    tail.pop();
                }
            }
            k--;
        }
        return res;
    }
};