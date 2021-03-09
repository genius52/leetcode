#include <string>
using namespace std;

class Solution_242 {
public:
    bool isAnagram(string s, string t) {
        int len1 = s.size();
        int len2 = t.size();
        if(len1 != len2)
            return false;
        int cnt1[26] = {0};
        int cnt2[26] = {0};
        for(int i = 0;i < len1;i++){
            cnt1[s[i] - 'a']++;
            cnt2[t[i] - 'a']++;
        }
        for(int i = 0;i < 26;i++){
            if(cnt1[i] != cnt2[i])
                return false;
        }
        return true;
    }
};