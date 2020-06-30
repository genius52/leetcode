#include <vector>
#include <algorithm>
#include <math.h>
using namespace std;

class Solution_452 {
public:
    int findMinArrowShots(vector<vector<int>>& points) {
        int len = points.size();
        if(len <= 1)
            return len;
        auto my_sort = [](const std::vector<int>& v1,const std::vector<int>& v2){
            if(v1[0] == v2[0]){
                return v1[1] < v2[1];
            }
            return v1[0] < v2[0];
        };
        std::sort(points.begin(),points.end(),my_sort);
        int cnt = 1;
        int pre_shot_pos = points[0][1];

        for(int i = 1;i < len;i++){
            if(points[i][0] <= pre_shot_pos){
                pre_shot_pos = min(points[i][1],pre_shot_pos);
                continue;
            }else{
                cnt++;
                pre_shot_pos = points[i][1];
            }
        }
        return cnt;
    }
};