#include <string>
using namespace std;

class Solution_5537 {
public:
    bool is_checkPalindromeFormation(string& s1) {
        int len = s1.size();
        if (len <= 1)
            return true;
        for (int i = 0; i < len / 2; i++) {
            if (s1[i] != s1[len - 1 - i])
                return false;
        }
        return true;
    }
    bool checkPalindromeFormation(string a, string b) {
        int len = a.size();
        if (len <= 1)
            return true;
        if(is_checkPalindromeFormation(a) || is_checkPalindromeFormation(b))
            return true;
        for (int i = 1; i < len; i++) {
            std::string prefix_a = a.substr(0,i);
            std::string suffix_a = a.substr(i, len - i);
            std::string prefix_b = b.substr(0, i);
            std::string suffix_b = b.substr(i, len - i);
            if (is_checkPalindromeFormation(prefix_a + suffix_b) || is_checkPalindromeFormation(prefix_b + suffix_a))
                return true;
        }
        return false;
    }
};