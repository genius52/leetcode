#include <math.h>
#include <string>

class Solution_1461 {
public:
    bool hasAllCodes(string s, int k) {
        //std::string sub(k, '0');
        int n = 1;
        int limit = 1 << k;
        std::string tag = "%0" + std::to_string(k) + "x";
        while (n < limit) {
            char* buf = new char[k];
            sprintf(buf, tag.c_str(), n);
            std::string sub;
            if (s.find(buf) == std::string::npos) {
                return false;
            }
            n++;
        }
        return true;
    }
};