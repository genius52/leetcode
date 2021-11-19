#include <vector>
#include <queue>
using namespace std;

//在 Δh 最大的那 ll 次使用梯子，而在剩余的情况下使用砖块
class Solution_1642 {
public:
    int furthestBuilding(vector<int>& heights, int bricks, int ladders) {
        std::priority_queue<int, std::vector<int>, std::greater<int> > q;//小顶堆
        int len = heights.size();
        int res = 0;
        for (int i = 0; i < len - 1; ++i) {
            if(heights[i] >=  heights[i + 1])
                continue;
            int diff = heights[i + 1] - heights[i];
            q.push(diff);
            if(q.size() > ladders){//不得不用砖块代替之前用的梯子
                auto top = q.top();
                bricks -= q.top();
                q.pop();//把diff最少的梯子换成砖块
            }
            if(bricks < 0){
                return i;
            }
        }
        return len - 1;
    }
};