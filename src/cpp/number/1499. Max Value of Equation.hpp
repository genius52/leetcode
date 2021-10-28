#include <vector>
#include <queue>
using namespace std;

//输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
//输出：4
//解释：前两个点满足 |xi - xj| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。
//第三个和第四个点也满足条件，得到值 10 + -10 + |5 - 6| = 1 。
//没有其他满足条件的点，所以返回 4 和 1 中最大的那个。

//yi - xi + yj + xj
class Solution_1499 {
public:
    //priority queue solution
    int findMaxValueOfEquation(vector<vector<int>>& points, int k) {
        int len = points.size();
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::less<std::pair<int,int>>> q;//xi,yi - xi
        int res = -2147483648;
        for(int i = 0;i < len;i++){
            while (!q.empty() && points[i][0] - q.top().second > k){ //只看最大的，删除后队列内可能还有不符合的xi，暂时不用管
                q.pop();
            }
            int cur_sum = points[i][0] + points[i][1];
            if(!q.empty()){
                res = max(res,cur_sum + q.top().first);
            }
            int cur_diff = points[i][1] - points[i][0];
            q.push({cur_diff,points[i][0]});
        }
        return res;
    }
    //deque solution
//    int findMaxValueOfEquation(vector<vector<int>>& points, int k) {
//        int len = points.size();
//        std::deque<std::pair<int,int>> q;//xi,yi - xi
//        int res = -2147483648;
//        for(int i = 0;i < len;i++){
//            int cur = points[i][0] + points[i][1];
//            while (!q.empty() && (points[i][0] - q.front().first) > k){ //删除不符合的xi
//                q.pop_front();
//            }
//            if(!q.empty()){
//                res = max(res,cur + q.front().second);
//            }
//            int cur_diff = points[i][1] - points[i][0];
//            while (!q.empty() && q.back().second <= cur_diff){//删除比当前yi - xi小的
//                q.pop_back();
//            }
//            q.push_back({points[i][0],cur_diff});
//        }
//        return res;
//    }
};