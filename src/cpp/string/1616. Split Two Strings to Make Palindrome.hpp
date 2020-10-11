#include <string>
using namespace std;

class Solution_1616 {
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

        int i = 0;
        for(;i < len/2;i++){
            if(a[i] != b[len -  1 - i]){
                break;
            }
        }
        {
            std::string prefix_a1 = a.substr(0,i);
            std::string prefix_a2 = a.substr(0,len - i);
            std::string suffix_a1 = a.substr(i, len - i);
            std::string suffix_a2 = a.substr(len - i,i);
            std::string prefix_b1 = b.substr(0, i);
            std::string prefix_b2 = b.substr(0,len - i);
            std::string suffix_b1 = b.substr(i, len - i);
            std::string suffix_b2 = b.substr(len - i, i);
            std::string sa1 = prefix_a1 + suffix_b1;
            std::string sa2 = prefix_a2 + suffix_b2;
            std::string sb1 = prefix_b1 + suffix_a1;
            std::string sb2 = prefix_b2 + suffix_a2;
            if(is_checkPalindromeFormation(sa1) || is_checkPalindromeFormation(sb1) ||
            is_checkPalindromeFormation(sa2) || is_checkPalindromeFormation(sb2))
                return true;
        }

        int j = 0;
        for (; j < len/2; ++j) {
            if(b[j] != a[len - 1 - j]){
                break;
            }
        }
        {
            std::string prefix_a1 = a.substr(0,j);
            std::string prefix_a2 = a.substr(0,len - j);
            std::string suffix_a1 = a.substr(j, len - j);
            std::string suffix_a2 = a.substr(len - j,j);
            std::string prefix_b1 = b.substr(0, j);
            std::string prefix_b2 = b.substr(0,len - j);
            std::string suffix_b1 = b.substr(j, len - j);
            std::string suffix_b2 = b.substr(len - j, j);
            std::string sa1 = prefix_a1 + suffix_b1;
            std::string sa2 = prefix_a2 + suffix_b2;
            std::string sb1 = prefix_b1 + suffix_a1;
            std::string sb2 = prefix_b2 + suffix_a2;
            if(is_checkPalindromeFormation(sa1) || is_checkPalindromeFormation(sb1) ||
               is_checkPalindromeFormation(sa2) || is_checkPalindromeFormation(sb2))
                return true;
        }
        return false;
    }
};