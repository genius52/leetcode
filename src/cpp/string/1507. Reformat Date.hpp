#include <string>
using namespace std;

//Input: date = "20th Oct 2052"
//Output: "2052-10-20"

//{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
class Solution_1507 {
public:
    string reformatDate(string date) {
        int len = date.size();
        int pos1 = 0;
        std::string day;
        while(pos1 < len && date[pos1] != ' '){
            if(day.empty() && (date[pos1] < '0' || date[pos1] > '9'))
                day = date.substr(0,pos1);
            pos1++;
        }
        pos1++;
        int pos2 = pos1;
        while(pos2 < len && date[pos2] != ' '){
            pos2++;
        }
        std::string month = date.substr(pos1,pos2 - pos1);
        pos2++;
        std::string year = date.substr(pos2,len - pos2);
        std::unordered_map<string,string> m;
        m["Jan"] = "01";
        m["Feb"] = "02";
        m["Mar"] = "03";
        m["Apr"] = "04";
        m["May"] = "05";
        m["Jun"] = "06";
        m["Jul"] = "07";
        m["Aug"] = "08";
        m["Sep"] = "09";
        m["Oct"] = "10";
        m["Nov"] = "11";
        m["Dec"] = "12";
        if(day.length() == 1)
            day = "0" + day;
        return year + "-" + m[month] + "-" + day;
    }
};