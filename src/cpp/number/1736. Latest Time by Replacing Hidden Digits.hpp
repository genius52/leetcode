#include <string>
using namespace std;

class Solution_1736 {
public:
    string maximumTime(string time) {
        std::string max_hour;
        std::string max_minute;
        if(time[0] == '?'){
            if(time[1] == '?'){
                max_hour = "23";
            }else if((time[1] - '0') >= 4){
                max_hour = "1";
                max_hour.push_back(time[1]);
            }else{
                max_hour = "2";
                max_hour.push_back(time[1]);
            }
        }else{
            if(time[1] == '?'){
                if(time[0] == '2'){
                    max_hour = "23";
                }else{
                    max_hour.push_back(time[0]);
                    max_hour += "9";
                }
            }else{
                max_hour.push_back(time[0]);
                max_hour.push_back(time[1]);
            }
        }
        if(time[3] == '?'){
            if(time[4] == '?'){
                max_minute = "59";
            }else{
                max_minute = "5";
                max_minute.push_back(time[4]);
            }
        }else{
            if(time[4] == '?'){
                max_minute.push_back(time[3]);
                max_minute += "9";
            }else{
                max_minute.push_back(time[3]);
                max_minute.push_back(time[4]);
            }
        }
        return max_hour + ":" + max_minute;
    }
};