#include <vector>
#include <string>
#include <map>
using namespace std;

class Solution_1629 {
public:
    char slowestKey(vector<int>& releaseTimes, string keysPressed) {
        std::map<char, int> record;
        int len = releaseTimes.size();
        record[keysPressed[0]] = releaseTimes[0];
        int max_duration = releaseTimes[0];
        for (int i = 1; i < len; i++) {
            if (record.find(keysPressed[i]) == record.end()) {
                record[keysPressed[i]] = releaseTimes[i] - releaseTimes[i - 1];
            }
            else {
                if (record[keysPressed[i]] < releaseTimes[i] - releaseTimes[i - 1])
                    record[keysPressed[i]] = releaseTimes[i] - releaseTimes[i - 1];
            }
            if (record[keysPressed[i]] > max_duration) {
                max_duration = record[keysPressed[i]];
                //res = keysPressed[i];
            }
        }
        char res = 0;
        for (auto it = record.rbegin(); it != record.rend(); it++) {
            if (it->second == max_duration) {
                res = it->first;
                break;
            }
        }
        return res;
    }
};