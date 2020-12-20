#include <vector>
#include <map>
#include <algorithm>
using namespace std;

//Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
//Output: 3
//Explanation:
//We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
//Then, we can reconstruct the sporting event as follows:
//We cut [1,9] into segments [1,2] + [2,8] + [8,9].
//Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].
bool my_sort(const std::vector<int>& v1, const vector<int>& v2)
{
    if(v1[0] != v2[0]){
        return v1[0] < v2[0];
    }else{
        return v1[1] < v2[1];
    }
}

class Solution_1024 {
public:
    int videoStitching(vector<vector<int>>& clips, int T) {
        std::sort(clips.begin(),clips.end(),my_sort);

        int len = clips.size();
        int start = T;
        int end = 0;
        int res = 0;
        for(int i = 0;i < len;i++){
            if(clips[i][0] != 0)
                break;
            int j = i + 1;
            int cur_cnt = 0;
            start = clips[i][0];
            end = clips[i][1];
            while(j < len && end < T){
                if (clips[j][1] > end){
                    end = clips[j][1];
                    cur_cnt++;
                }

            }
        }
        return res;
    }
};