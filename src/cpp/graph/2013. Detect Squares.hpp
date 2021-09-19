#include <vector>
#include <unordered_map>

class DetectSquares {
    std::unordered_map<int, int> record;
public:
    DetectSquares() {

    }

    void add(vector<int> point) {
        int key = point[0] * 10000 + point[1];
        record[key]++;
    }

    int count(vector<int> point) {
        int res = 0;
        int x1 = point[0];
        int y1 = point[1];
        int key = x1 * 10000 + y1;
        int cnt1 = 1;
        if (record.find(key) != record.end()) {
            cnt1 = record[key];
        }
        for (auto it = record.begin(); it != record.end(); it++) {
            int key2 = (*it).first;
            int cnt2 = (*it).second;
            int x2 = key2 / 10000;
            int y2 = key2 % 10000;
            if (x1 == x2 || y1 == y2 || abs(x1 - x2) != abs(y1 - y2)) {
                continue;
            }
            int x3 = x1;
            int y3 = y2;
            int key3 = x3 * 10000 + y3;
            int x4 = x2;
            int y4 = y1;
            int key4 = x4 * 10000 + y4;
            if (record.find(key3) != record.end() && record.find(key4) != record.end()) {
                int cnt3 = record[key3];
                int cnt4 = record[key4];
                res += cnt2 * cnt3 * cnt4;
            }
        }
        return res;
    }
};