#include <string>
using namespace std;

class Solution_925 {
public:
    bool isLongPressedName(string name, string typed) {
        int len = name.length();
        int len_t = typed.length();
        if(len > len_t)
            return false;
        int i = 0;
        int j = 0;
        while(i < len || j < len_t){
            char prev = name[i];
            if(name[i] != typed[j])
                return false;
            i++;
            j++;
            if(name[i] != typed[j]){
                while(prev == typed[j] && j < len_t){
                    j++;
                }
            }
        }
        return true;
    }
};