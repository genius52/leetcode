#include <string>
using namespace std;

class Solution_125 {
public:
    bool isPalindrome(string s) {
        if(s.empty())
            return true;
        int len = s.length();
        int low = 0;
        int high = len - 1;
        while(low <= high){
            if((s[low] >= '0' && s[low] <= '9') || (s[high] >= '0' && s[high] <= '9')){
                if(s[low] != s[high])
                    return false;
                else
                {
                    low++;
                    high--;
                    continue;
                }
            }
            if(s[low] < 'A' || s[low] > 'z'){
                low++;
                continue;
            }
            if(s[high] < 'A' || s[high] > 'z'){
                high--;
                continue;
            }
            int diff = abs(s[low]-s[high]);
            if(diff != 0 && diff != 32)
                return false;
            low++;
            high--;
        }
        return true;
    }
};