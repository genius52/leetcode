#include <string>
using namespace std;

class Solution_459 {
public:
    bool repeatedSubstringPattern(string s) {
        int len = s.length();
        if(len <= 1)
            return false;
        for(int i = 0 ;i < len/2;i++){
            int sub_len = i + 1;
            string sub = s.substr(0,sub_len);
            bool match = true;
            for(int j = i + 1;j < len ;j = j + sub_len){
                if(s.find(sub,j) != j){
                    match = false;
                    break;
                }
            }
            if(match)
                return true;
        }
        return false;
    }
};