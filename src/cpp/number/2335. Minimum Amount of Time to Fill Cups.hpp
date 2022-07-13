#include <vector>
#include <queue>

class Solution_2335 {
public:
    int fillCups(vector<int>& amount) {
        std::priority_queue<int, std::vector<int>, std::less<int> > q;
        if(amount[0] != 0)
            q.push(amount[0]);
        if(amount[1] != 0)
            q.push(amount[1]);
        if(amount[2] != 0)
            q.push(amount[2]);
        int res = 0;
        while (q.size() > 1){
            int top1 = q.top();
            q.pop();
            int top2 = q.top();
            q.pop();
            int top3 = q.top();
            int gap = top2 - top3 + 1;
            res += gap;
            top1 -= gap;
            top2 -= gap;
            if(top1 != 0)
                q.push(top1);
            if(top2 != 0)
                q.push(top2);
        }
        if(!q.empty())
            res += q.top();
        return res;
    }
};