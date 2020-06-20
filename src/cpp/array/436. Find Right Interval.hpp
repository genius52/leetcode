#include <vector>
#include <map>
using namespace std;

//Input: [ [3,4], [2,3], [1,2] ]
//
//Output: [-1, 0, 1]
class Solution_436 {
public:
    vector<int> findRightInterval(vector<vector<int>>& intervals) {
        int len = intervals.size();
        std::vector<int> res;
        res.resize(len);
        std::map<int,int> record;
        for(int i = 0;i < len;i++){
            record[intervals[i][0]] = i;
        }
        for(int i = 0;i < len;i++){
            auto it = record.lower_bound(intervals[i][1]);
            if(it != record.end())
                res[i] = (*it).second;
            else
                res[i] = -1;
        }
        return res;
    }
};