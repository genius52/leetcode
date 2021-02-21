#include <vector>
#include <string>
using namespace std;

class Solution_1769 {
public:
    vector<int> minOperations(string boxes) {
        int len = boxes.size();
        std::vector<int> res(len);
        for (int i = 0; i < len; i++) {
            int steps = 0;
            for (int j = 0; j < len; j++) {
                if (i == j)
                    continue;
                if (boxes[j] == '1') {
                    steps += abs(i - j);
                }
            }
            res[i] = steps;
        }
        return res;
    }
};
