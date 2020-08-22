#include <string>
#include <unordered_map>
#include <math.h>
using namespace std;

class Solution_166 {
public:
    string fractionToDecimal(int numerator, int denominator) {
        if(numerator < 0 && denominator < 0){
            numerator = - numerator;
            denominator = -denominator;
        }
        std::string res;

        while(abs(numerator) >= abs(denominator)){
            int interger = numerator / denominator;
            res += std::to_string(interger);
            numerator = numerator %  denominator;
        }
        if(res.size() == 0)
            res += "0";

        std::unordered_map<int,int> record;
        numerator = abs(numerator - numerator / denominator);
        denominator = abs(denominator);
        if(numerator > 0)
            res += ".";
        int repeat_num = -1;
        while(numerator > 0){
            numerator = numerator * 10;

            if(record.find(numerator/denominator) != record.end()){
                repeat_num = numerator/denominator;
                break;
            }
            res += std::to_string(numerator / denominator);
            record[numerator / denominator]++;
            numerator = numerator % denominator;
        }
        if(repeat_num != -1){
            int pos = res.rfind(std::to_string(repeat_num));
            res = res.substr(0,pos) + "(" + res.substr(pos,res.size() - pos) + ")";
        }
        return res;
    }
};