#include <vector>
using namespace std;

class Solution_1921 {
public:
    int eliminateMaximum(vector<int>& dist, vector<int>& speed) {
        int len = dist.size();
        std::vector<float> record(len);
        for (int i = 0; i < len; i++) {
            record[i] = float(dist[i]) / float(speed[i]);
        }
        std::sort(record.begin(), record.end());
        int timestamp = 0;
        int res = 0;
        for (int i = 0; i < len; i++) {
            if (record[i] > timestamp) {
                timestamp++;
                res++;
            }
            else {
                break;
            }
        }
        return res;
    }
};