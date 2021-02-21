#include <string>
using namespace std;

class Solution_1768 {
public:
    string mergeAlternately(string word1, string word2) {
        int len1 = word1.size();
        int len2 = word2.size();
        int l1 = 0;
        int l2 = 0;
        std::string res;
        while (l1 < len1 && l2 < len2) {
            res += word1[l1++];
            res += word2[l2++];
        }
        while (l1 < len1) {
            res += word1[l1++];
        }
        while (l2 < len2) {
            res += word2[l2++];
        }
        return res;
    }
};
