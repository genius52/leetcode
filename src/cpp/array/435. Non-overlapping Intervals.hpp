#include <vector>
#include <unordered_map>
using namespace std;

//Input: [[1,2],[2,3],[3,4],[1,3]]
//Output: 1
//Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
bool comp(const std::vector<int>& v1,const std::vector<int>& v2)
{
    return v1[1] < v2[1];
}

class Solution_435 {
public:
    int eraseOverlapIntervals(vector<vector<int>>& intervals) {
        int len = intervals.size();
        if(len <= 1)
            return 0;
        auto mycmp = [](const std::vector<int>& v1,const std::vector<int>& v2){
            return v1[1] < v2[1];
        };
        std::sort(intervals.begin(),intervals.end(),mycmp);
        int nodup_cnt = 1;
        int last_end = intervals[0][1];
        for(int i = 1;i < len;i++){
            if(intervals[i][0] >= last_end){
                nodup_cnt++;
                last_end = intervals[i][1];
            }
        }
        return len - nodup_cnt;
    }
};