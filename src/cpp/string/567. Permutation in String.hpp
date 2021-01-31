#include <string>
#include <set>
using namespace std;

class Solution_567 {
public:
    bool checkInclusion(string s1, string s2) {
        int len1 = s1.length();
        int len2 = s2.length();
        if(len1 > len2){
            return false;
        }
        std::multiset<char> set1(s1.begin(),s1.end());
        int start = 0;
        std::multiset<char> set2(s2.begin(),s2.begin() + len1);
        while(start <= len2 - len1){
            if(set2 == set1){
                return true;
            }
            if(start == len2 - len1)
                break;
            set2.erase(set2.find(s2[start]));
            start++;
            set2.insert(s2[start + len1 - 1]);
        }
        return false;
    }
};