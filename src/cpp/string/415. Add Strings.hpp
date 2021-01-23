#include <string>
using namespace std;

class Solution {
public:
    string addStrings(string num1, string num2) {
        string sum;
        int len1 = num1.length();
        if(len1 == 0)
            return num2;
        int len2 = num2.length();
        if(len2 == 0)
            return num1;
        string result;
        int min_len = len1>len2?len2:len1;
        int index = 0;
        int upgrade = 0;
        while(index < min_len){
            int v1 = num1.at(len1 - 1 - index) - '0';
            int v2 = num2.at(len2 - 1 - index) - '0';
            int sum = v1 + v2 + upgrade;
            if(sum >=10)
            {
                result = std::to_string(sum - 10) + result;
                upgrade = 1;
            }
            else{
                result = std::to_string(sum) + result;
                upgrade = 0;
            }
            index++;
        }
        while(index < len1)
        {
            int val = num1.at(len1 - 1 - index) - '0';
            int sum = val + upgrade;
            if(sum >= 10){
                result = std::to_string(sum - 10) + result;
                upgrade = 1;
            }
            else{
                result = std::to_string(sum) + result;
                upgrade = 0;
            }
            index++;
        }
        while(index < len2)
        {
            int val = num2.at(len2 - 1 - index) - '0';
            int sum = val + upgrade;
            if(sum >= 10){
                result = std::to_string(sum - 10) + result;
                upgrade = 1;
            }
            else{
                result = std::to_string(sum) + result;
                upgrade = 0;
            }
            index++;
        }
        if(upgrade == 1)
        {
            result = "1" + result;
        }
        return result;
    }
};