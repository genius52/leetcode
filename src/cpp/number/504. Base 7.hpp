#include <string>

//Input: 100
//Output: "202"
class Solution_504 {
public:
    string convertToBase7(int num) {
        if(num == 0)
            return "0";
        bool is_negtive = num < 0;
        num = abs(num);
        std::string res;
        while(num > 0){
            int mod = num % 7;
            res = std::to_string(mod) + res;
            num = num / 7;
        }
        if(is_negtive)
            res = "-" + res;
        return res;
    }
};