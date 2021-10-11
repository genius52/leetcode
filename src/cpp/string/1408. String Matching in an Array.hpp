#include <vector>
#include <string>
using namespace std;

class Solution_1408 {
public:
    vector<string> stringMatching(vector<string>& words) {
        int len = words.size();
        std::vector<string> res;
        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (i == j)
                    continue;
                if (words[j].find(words[i]) != std::string::npos) {
                    res.push_back(words[i]);
                    break;
                }
            }
        }
        return res;
    }
};