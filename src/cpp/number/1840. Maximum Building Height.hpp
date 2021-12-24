#include <vector>
#include <map>
using namespace std;

class Solution_1840 {
public:
    int maxBuilding(int n, vector<vector<int>>& restrictions) {
        std::sort(restrictions.begin(),restrictions.end());
        int len = restrictions.size();
        int pre_height = 0;
        int pre_idx = 1;
        for(int i = 0;i < len - 1;i++){
            int cur_height = restrictions[i][1];
            if(restrictions[i][1] > pre_height){
                cur_height = min(cur_height,pre_height + restrictions[i][0] - pre_idx);
            }
            if(restrictions[i][1] > restrictions[i + 1][1]){
                cur_height = min(cur_height,restrictions[i + 1][1] + restrictions[i + 1][0] - restrictions[i][0]);
            }
            restrictions[i][1] = cur_height;
            pre_height = cur_height;
            pre_idx = restrictions[i][0];
        }
        if(len > 0){
            int last_height = restrictions[len - 1][1];
            if(last_height > pre_height){
                last_height = min(last_height,pre_height + restrictions[len - 1][0] - pre_idx);
            }
            restrictions[len - 1][1] = last_height;
        }
        //FROM END TO START
        if(len > 0){
            pre_height = restrictions[len - 1][1];
            pre_idx = restrictions[len - 1][0];
            for(int i = len - 2;i >= 0;i--){
                int cur_height = restrictions[i][1];
                if(restrictions[i][1] > pre_height){
                    cur_height = min(cur_height,pre_height + pre_idx - restrictions[i][0]);
                }
                if(i > 0){
                    if(restrictions[i][1] > restrictions[i - 1][1]){
                        cur_height = min(cur_height,restrictions[i - 1][1] + restrictions[i][0] - restrictions[i - 1][0]);
                    }
                }
                restrictions[i][1] = cur_height;
                pre_height = cur_height;
                pre_idx = restrictions[i][0];
            }
            int first_height = restrictions[0][1];
            if(first_height > pre_height){
                first_height = min(first_height,pre_height + pre_idx - restrictions[0][0]);
            }
            restrictions[0][1] = first_height;
        }
        pre_height = 0;
        pre_idx = 1;
        int max_height = 0;
        for(int i = 0;i < len;i++){
            int height_diff = abs(restrictions[i][1] - pre_height);
            int distance = restrictions[i][0] - pre_idx;
            if(height_diff < distance){
                int cur_max = (distance - height_diff ) /2 + max(restrictions[i][1],pre_height);
                max_height = max(max_height,cur_max);
            }else{
                max_height = max(max_height,pre_height);
                max_height = max(max_height,restrictions[i][1]);
            }
            pre_height = restrictions[i][1];
            pre_idx = restrictions[i][0];
        }
        int last_height = pre_height + n - pre_idx;
        max_height = max(max_height,last_height);
        return max_height;
    }
};