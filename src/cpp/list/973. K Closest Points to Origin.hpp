#include <vector>
#include <queue>
using namespace std;

class Solution_973 {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        std::vector<std::vector<int>> res(k);
        std::priority_queue<std::pair<int,int>> q;
        int len = points.size();
        for(int i = 0;i < len;i++){
            if(q.size() < k){
                q.push({points[i][0] * points[i][0] + points[i][1] * points[i][1],i});
            }else{
                int dis = points[i][0] * points[i][0] + points[i][1] * points[i][1];
                if (dis < q.top().first){
                    q.pop();
                    q.push({dis,i});
                }
            }
        }
        for (int i = 0;i < k;i++){
            res[i] = points[q.top().second];
            q.pop();
        }
        return res;
    }
};