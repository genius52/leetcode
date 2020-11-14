#include <string>
#include <unordered_map>
using namespace std;

class Solution_1624 {
public:
    int maxLengthBetweenEqualCharacters(string s) {
        std::unordered_map<char, int> begin_record;
        std::unordered_map<char, int> end_record;
        int len = s.size();
        int res = -1;
        for (int i = 0; i < len; i++) {
            if (begin_record.find(s[i]) == begin_record.end()) {
                begin_record[s[i]] = i;
            }
            else {
                end_record[s[i]] = i;
                int maxlen = i - begin_record[s[i]] - 1;
                if (maxlen > res) {
                    res = maxlen;
                }
            }
        }
        return res;
    }
};