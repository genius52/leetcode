#include <string>
#include <vector>
using namespace std;

class Solution_953 {
public:
    bool cmp(std::string& s1,std::string& s2,std::string& pattern){
        int len1 = s1.length();
        int len2 = s2.length();
        int min_len = len1 > len2?len2:len1;
        for(int i = 0;i < min_len;i++){
            auto c1 = s1[i];
            auto c2 = s2[i];
            auto pos1 = pattern.find(c1);
            auto pos2 = pattern.find(c2);
            if(pos1 > pos2)
                return false;
            else if(pos1 < pos2)
                return true;
        }
        if(len1 > len2)
            return false;
        return true;
    }
    bool isAlienSorted(vector<string>& words, string order) {
        int len = words.size();
        if(len <= 1)
            return true;
        for(int i = 0;i < len - 1;i++){
            auto first = words[i];
            auto second = words[i+1];
            if(!cmp(first,second,order))
                return false;
        }
        return true;
    }
};