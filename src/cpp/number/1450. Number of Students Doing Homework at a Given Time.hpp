#include <vector>

class Solution_1450 {
public:
    int busyStudent(vector<int>& startTime, vector<int>& endTime, int queryTime) {
        int cnt = 0;
        int len = startTime.size();
        for(int i = 0;i < len;i++){
            if(startTime[i] <= queryTime && queryTime <= endTime[i])
                cnt++;
        }
        return cnt;
    }
};